package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API")

	client := openai.NewClient(apiKey)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a the game master. You are talking with the player of the game. Your job is to create a choose your own adventure, coupled with a specific musical challenge at each step of the journey. Create the adventure based on the fears and loves of the player. The musical challenges should be challenging, and should be centered around music theory. It should also be noted that this is a text based game, so the player would need to be able to answer the questions without any audio. You can determine if the player has answered the question correctly or not. If you determine that the player answered incorrectly, you may end the game. If the player successfully solves 3 musical puzzels, they win the game. Try to weave in the musical challenges to the plot of the story. The player should be asked alternative questions about which direction they'd like to go, or which action they'd like to take, then be asked the musical challenges based on their action or choice of direction. The player's name is Keegan, his worst fear is harm coming to his family, and his greatest loves are family, guitar, and programming. ",
			},
		},
	}

	fmt.Println("Musical Adventure")
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: s.Text(),
		})
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)
		req.Messages = append(req.Messages, resp.Choices[0].Message)
		fmt.Print("> ")
	}
}
