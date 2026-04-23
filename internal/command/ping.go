package command

import (
	"github.com/amatsagu/tempest"
)

var Ping = tempest.Command{
	Name:                "ping",
	Description:         "Ping!",
	SlashCommandHandler: handlePing,
}

func handlePing(itx *tempest.CommandInteraction) {
	itx.SendLinearReply("Pong! :ping_pong:", false)
}
