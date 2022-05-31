package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
)

var (
	token   = os.Getenv("disgo_token")
	guildID = snowflake.GetEnv("disgo_guild_id")

	commands = []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			CommandName: "modal",
			Description: "brings up a modal",
		},
	}
)

func main() {
	log.SetLevel(log.LevelTrace)
	log.Info("starting example...")
	log.Info("disgo version: ", disgo.Version)

	client, err := disgo.New(token,
		bot.WithGatewayConfigOpts(gateway.WithGatewayIntents(discord.GatewayIntentsNone)),
		bot.WithEventListenerFunc(commandListener),
		bot.WithEventListenerFunc(modalListener),
	)
	if err != nil {
		log.Fatal("error while building disgo instance: ", err)
		return
	}

	defer client.Close(context.TODO())

	if _, err = client.Rest().SetGuildCommands(client.ApplicationID(), guildID, commands); err != nil {
		log.Fatal("error while registering commands: ", err)
	}

	if err = client.ConnectGateway(context.TODO()); err != nil {
		log.Fatal("error while connecting to gateway: ", err)
	}

	log.Infof("example is now running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}

func commandListener(event *events.ApplicationCommandInteractionCreate) {
	data := event.SlashCommandInteractionData()
	if data.CommandName() == "modal" {
		if err := event.CreateModal(discord.NewModalCreateBuilder().
			SetTitle("Modal Title").
			SetCustomID("modal-id").
			AddActionRow(discord.NewShortTextInput("short-text-input", "short text")).
			AddActionRow(discord.NewParagraphTextInput("paragraph-text-input", "paragraph text")).
			AddActionRow(discord.NewSelectMenu("select-menu", "select something idiot",
				discord.NewSelectMenuOption("helo", "helo"),
				discord.NewSelectMenuOption("uwu", "uwu"),
				discord.NewSelectMenuOption("owo", "owo"),
			)).
			Build(),
		); err != nil {
			event.Client().Logger().Error("error creating modal: ", err)
		}
	}
}

func modalListener(event *events.ModalSubmitInteractionCreate) {
	var content string
	for customID, component := range event.Data.Components {
		switch c := component.(type) {
		case discord.TextInputComponent:
			content += customID.String() + ": " + c.Value + "\n"
		case discord.SelectMenuComponent:
			content += customID.String() + ": " + strings.Join(c.Values, ", ") + "\n"
		}
	}
	if err := event.CreateMessage(discord.MessageCreate{Content: content}); err != nil {
		event.Client().Logger().Error("error creating modal: ", err)
	}
}
