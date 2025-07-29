package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/xtt28/freakbot/internal/classifier"
)

func main() {
	godotenv.Load()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "must include text argument to classify")
		os.Exit(1)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "must specify the environment variable OPENAI_API_KEY")
		os.Exit(1)
	}

	service := classifier.NewOpenAIClassifier(apiKey)
	message := strings.Join(os.Args[1:], " ")

	result, err := service.IsFlagged(message)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("result: %t\n", result)
}
