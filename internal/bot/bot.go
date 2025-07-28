package bot

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/xtt28/freakbot/internal/classifier"
	"github.com/xtt28/freakbot/internal/handler"
	"github.com/xtt28/freakbot/internal/repository"
)

type BotApp struct {
	discordSess *discordgo.Session
	dbConn repository.Connection
	classifierService classifier.ClassifierService
	handler *handler.Handler
}

type BotAppParams struct {
	DiscordToken string
	OpenAIKey string
	DatabaseDSN string
}

func (b *BotApp) ready(s *discordgo.Session, event *discordgo.Ready) {
	b.handler = handler.NewHandler(b.dbConn, b.classifierService)
	b.discordSess.AddHandler(b.handler.GuildCreate)
	b.discordSess.AddHandler(b.handler.MessageCreate)

	log.Println("bot is ready")

	s.UpdateGameStatus(0, "freaky mode")
}

func (b *BotApp) Run() {
	b.discordSess.AddHandler(b.ready)

	b.discordSess.Open()
	defer b.discordSess.Close()

	log.Println("running bot")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func New(params BotAppParams) (*BotApp, error) {
	app := &BotApp{}

	conn, err := repository.NewGORMSQLiteConnection(params.DatabaseDSN)
	if err != nil {
		return nil, err
	}
	app.dbConn = conn

	sess, err := discordgo.New(params.DiscordToken)
	if err != nil {
		return nil, err
	}
	app.discordSess = sess

	serv := classifier.NewOpenAIClassifier(params.OpenAIKey)
	app.classifierService = serv

	return app, nil
}