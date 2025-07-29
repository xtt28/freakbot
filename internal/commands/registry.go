package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/xtt28/freakbot/internal/classifier"
	"github.com/xtt28/freakbot/internal/manifest"
	"github.com/xtt28/freakbot/internal/repository"
)

type CommandRegistry struct {
	discordSess       *discordgo.Session
	dbConn            repository.Connection
	classifierService classifier.ClassifierService
	commands          []*discordgo.ApplicationCommand
	handlerMap        map[string]func(*discordgo.Session, *discordgo.Interaction)
}

func (r *CommandRegistry) registerCommands() {
	log.Println("registering commands")

	r.commands = []*discordgo.ApplicationCommand{
		{
			Name:        "about",
			Description: "Shows you information about " + manifest.Name,
		},
		{
			Name:        "freakerboard",
			Description: "Shows you a list of the freakiest people in this server.",
		},
	}

	r.handlerMap = map[string]func(*discordgo.Session, *discordgo.Interaction){
		"about":        r.About,
		"freakerboard": r.Leaderboard,
	}

	for _, v := range r.commands {
		log.Printf("registering command %s\n", v.Name)
		_, err := r.discordSess.ApplicationCommandCreate(r.discordSess.State.User.ID, "", v)
		if err != nil {
			log.Println("could not register command", err)
		}
	}
}

func (r *CommandRegistry) HandleCommand(name string, s *discordgo.Session, i *discordgo.Interaction) {
	handler, ok := r.handlerMap[name]
	if !ok {
		return
	}

	log.Printf("handling command %s\n", name)
	handler(s, i)
}

func NewRegistry(
	discordSess *discordgo.Session,
	dbConn repository.Connection,
	classifierService classifier.ClassifierService,
) *CommandRegistry {
	r := &CommandRegistry{
		discordSess,
		dbConn,
		classifierService,
		nil,
		nil,
	}
	r.registerCommands()
	return r
}
