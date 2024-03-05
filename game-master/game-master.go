package gamemaster

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func Play() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API")

	client := openai.NewClient(apiKey)

	s := bufio.NewScanner(os.Stdin)

	var answers []string
	getUserInfo(s, &answers)

	if len(answers) < 3 {
		log.Fatal("Did not get all user info")
	}

	player := NewPlayer(answers[0], answers[1], answers[2])

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: fmt.Sprintf("You are a the game master. You are talking with the player of the game. Your job is to create a choose your own adventure, coupled with a specific musical challenge at each step of the journey. Create the adventure based on the fears and loves of the player. The musical challenges should be challenging, and should be centered around music theory. Example question: How do you spell a G major chord? Example: Spell the notes of an A major scale. It should also be noted that this is a text based game, so the player would need to be able to answer the questions without any audio. You can determine if the player has answered the question correctly or not. If you determine that the player answered incorrectly, you may end the game. If the player successfully solves 3 musical puzzels, they win the game. Try to weave in the musical challenges to the plot of the story. The player should be asked alternative questions about which direction they'd like to go, or which action they'd like to take, then be asked the musical challenges based on their action or choice of direction. Player name: %s, loves: %s, fears: %s ", player.Name, player.Loves, player.Fears),
			},
		},
	}

	fmt.Println("Musical Adventure")
	fmt.Println("Great the gamemaster to get started")
	fmt.Println("---------------------\n")
	fmt.Print("> ")
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

func getUserInfo(s *bufio.Scanner, ans *[]string) {
	var count int
	fmt.Println("What is your name?")
	fmt.Print("> ")
	for s.Scan() {
		*ans = append(*ans, s.Text())
		switch count {
		case 0:
			fmt.Println("What is/are your greatest loves?")
			fmt.Print("> ")
			count++
		case 1:
			fmt.Println("What is/are your greatest fears?")
			fmt.Print("> ")
			count++
		default:
			return
		}
	}
}
