package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/amatsagu/tempest"
)

var Server = tempest.Command{
	Name:        "server",
	Description: "Start or stop a game server.",
	Options: []tempest.CommandOption{
		{
			Name:        "game",
			Description: "Game server to start or stop",
			Type:        tempest.STRING_OPTION_TYPE,
			Required:    true,
			Choices: []tempest.Choice{
				{
					Name:  "Satisfactory",
					Value: "satisfactory",
				},
			},
		},
		{
			Name:        "action",
			Description: "Start or stop the server",
			Type:        tempest.STRING_OPTION_TYPE,
			Required:    true,
			Choices: []tempest.Choice{
				{
					Name:  "Start",
					Value: "start",
				},
				{
					Name:  "Stop",
					Value: "stop",
				},
			},
		},
	},
	SlashCommandHandler: handleServer,
}

type Workflow struct {
	Name       string
	Owner      string
	Repository string
	Actions    map[string]string
	Ref        string
}

var workflows = map[string]Workflow{
	"satisfactory": {
		Name:       "Satisfactory",
		Owner:      "paschi",
		Repository: "satisfactory-hcloud",
		Actions: map[string]string{
			"start": "deploy.yml",
			"stop":  "destroy.yml",
		},
		Ref: "main",
	},
}

func handleServer(itx *tempest.CommandInteraction) {
	game, _ := itx.GetOptionValue("game")
	action, _ := itx.GetOptionValue("action")
	workflow := workflows[game.(string)]
	itx.SendLinearReply(fmt.Sprintf("<@%d> wants to **%s** the **%s** server. I'm working on it!", itx.Member.User.ID, action, workflow.Name), false)
	itx.Defer(false)
	err := triggerWorkflow(workflow, action.(string))
	if err != nil {
		itx.SendFollowUp(tempest.ResponseMessageData{
			Content: fmt.Sprintf(":warning: **Failed** to **%s** the **%s** server: `%s`", action, workflow.Name, err.Error()),
		}, false)
		return
	}
	channelID, err := tempest.StringToSnowflake(os.Getenv("DISCORD_STATUS_CHANNEL_ID"))
	itx.SendFollowUp(tempest.ResponseMessageData{
		Content: fmt.Sprintf(":white_check_mark: Workflow successfully triggered! Please check <#%d> for updates!", channelID),
	}, false)
}

func triggerWorkflow(workflow Workflow, action string) error {
	token := os.Getenv("GITHUB_ACCESS_TOKEN")
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows/%s/dispatches", workflow.Owner, workflow.Repository, workflow.Actions[action])
	payload, err := json.Marshal(map[string]any{
		"ref": workflow.Ref,
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to trigger workflow, status: %s", resp.Status)
	}
	return nil
}
