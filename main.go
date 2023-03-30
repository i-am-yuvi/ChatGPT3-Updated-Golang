package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func GetResponse(client gpt3.Client, ctx context.Context, chat string) {

	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			chat,
		},
		MaxTokens:   gpt3.IntPtr(30),
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		fmt.Print(resp.Choices[0].Text)
	})
	if err != nil {
		log.Fatalln(err)
		os.Exit(13)
	}
	fmt.Printf("\n")
}

// type Nullwriter int

// func (Nullwriter) Write([]byte) (int, error) { return 0, nil }

func main() {
	// Load .env file
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API_KEY missing from .env file")
	}

	// Create a new GPT-3 client
	client := gpt3.NewClient(apiKey)
	ctx := context.Background()

	rootCmd := &cobra.Command{

		Use:   "chatgpt",
		Short: "Chat with GPT-3",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Println("Say something('quit' to end)")
				if !scanner.Scan() {
					break
				}

				chat := scanner.Text()
				switch chat {
				case "quit":
					quit = true

				default:
					GetResponse(client, ctx, chat)
				}
			}

		},
	}

	rootCmd.Execute()

}
