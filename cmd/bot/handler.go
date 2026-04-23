package bot

import (
	"log"
	"net/http"
	"os"

	"github.com/amatsagu/tempest"
	"github.com/paschi/jackbot/internal/command"
)

var client tempest.Client

func syncCommands() {
	guildID, err := tempest.StringToSnowflake(os.Getenv("DISCORD_GUILD_ID"))
	if err != nil {
		log.Fatalln("failed to parse 'DISCORD_GUILD_ID' env variable to snowflake", err)
	}
	err = client.SyncCommandsWithDiscord([]tempest.Snowflake{guildID}, nil, false)
	if err != nil {
		log.Fatalln("failed to sync local commands storage with Discord API", err)
	}
}

func init() {
	client = tempest.NewClient(tempest.ClientOptions{
		Token:     os.Getenv("DISCORD_BOT_TOKEN"),
		PublicKey: os.Getenv("DISCORD_PUBLIC_KEY"),
	})
	client.RegisterCommand(command.Ping)
	client.RegisterCommand(command.Roll)
	client.RegisterCommand(command.Server)
	if os.Getenv("SYNCHRONIZE_COMMANDS") == "true" {
		syncCommands()
	}
}

func Handle(w http.ResponseWriter, r *http.Request) {
	client.DiscordRequestHandler(w, r)
}
