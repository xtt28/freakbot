package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/xtt28/freakbot/internal/bot"
)

func main() {
	godotenv.Load()

	b, err := bot.New(bot.BotAppParams{
		DiscordToken: os.Getenv("DISCORD_TOKEN"),
		OpenAIKey:    os.Getenv("OPENAI_API_KEY"),
		DatabaseDSN:  os.Getenv("DB_DSN"),
	})

	if err != nil {
		log.Fatalln("could not initialize bot application instance", err)
		return
	}

	b.Run()
}
