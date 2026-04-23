package command

import (
	"fmt"
	"math/rand"

	"github.com/amatsagu/tempest"
)

var Roll = tempest.Command{
	Name:        "roll",
	Description: "Roll the dice!",
	Options: []tempest.CommandOption{
		{
			Name:        "sides",
			Description: "Number of sides on the dice",
			Type:        tempest.INTEGER_OPTION_TYPE,
			Required:    true,
			MinValue:    2,
			MaxValue:    100,
		},
	},
	SlashCommandHandler: handleRoll,
}

func handleRoll(itx *tempest.CommandInteraction) {
	sides, _ := itx.GetOptionValue("sides")
	sidesInt := int(sides.(float64))
	result := rand.Intn(sidesInt) + 1
	itx.SendLinearReply(fmt.Sprintf("<@%d> rolled a **%d** (1-%d)! :game_die:", itx.Member.User.ID, result, sidesInt), false)
}
