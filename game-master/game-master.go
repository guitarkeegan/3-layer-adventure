package gamemaster

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

func Play() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API")

	client := openai.NewClient(apiKey)
	ctx := context.Background()
	// weird
	ff := new(FunkFactory)

	s := bufio.NewScanner(os.Stdin)

	// maybe array
	//
	var answers []string
	// return user info
	// newUserInfo(os.Stdin)
	// s := bufio.NewScanner(os.Stdin)
	// 1. call getUserInfo and pass in os.Stdin
	// 2. return UserInfo (name, loves, fears) or error
	// 3. newPlayer(UserInfo)
	//
	// for holding scope, use a pointer
	// for data, use 'make' and use and instance rather than a pointer
	getUserInfo(s, &answers)

	if len(answers) < 3 {
		log.Fatal("Did not get all user info")
	}

	player := NewPlayer(answers[0], answers[1], answers[2])

	endGameParams := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"prompt": {
				Type:        jsonschema.String,
				Description: "This is a prompt for an image generation. If the user has won the game, they should see a celebratory image that includes things that they love. If they have lost, they should see the things that they fear. Adjust this image prompt according to whether or not the user won the game. Also include relavent information from the adventure that they just had in the prompt",
			},
		},
		Required: []string{"prompt"},
	}

	endGameFunc := openai.FunctionDefinition{
		Name:        "endGame",
		Description: "generate a prompt, that will be used to generate an image for he user, based whether they won or lost",
		Parameters:  endGameParams,
	}

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo0125,
		Tools: []openai.Tool{
			{Type: openai.ToolTypeFunction, Function: &endGameFunc},
		},
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: fmt.Sprintf("You are a the game master. You are talking with the player of the game. Your job is to create a choose your own adventure, coupled with a specific musical challenge at each step of the journey. Create the adventure based on the fears and loves of the player. The musical challenges should be challenging, and should be centered around music theory. Example question: How do you spell a G major chord? Example: Spell the notes of an A major scale. It should also be noted that this is a text based game, so the player would need to be able to answer the questions without any audio. You determine if the player has answered the question correctly or not. If you determine that the player answered incorrectly, you may end the game and generate an image for the user by calling the endGame function. If the player successfully solves 3 musical puzzels, they win the game, and you call the endGame function to show the image to the user. If they answer any one musical question incorrectly, they lose the game. Try to weave in the musical challenges to the plot of the story. The player should be asked alternative questions about which direction they'd like to go, or which action they'd like to take, then be asked the musical challenges based on their action or choice of direction. Player name: %s, loves: %s, fears: %s ", player.Name, player.Loves, player.Fears),
			},
		},
	}

	fmt.Println("Musical Adventure")
	fmt.Println("Greet the gamemaster to get started")
	fmt.Println("---------------------\n")
	fmt.Print("> ")
	for s.Scan() {
		// move to gameai package
		// gameai type will save state
		// field called state (unexported)
		// expose state to user with a method called State()
		// constants - exported, that represent state
		// gameai.New() state will be chatting
		// gai.Chat() state will be 'ending' 'chatting' ?
		// if in 'ending' can't Chat() would hit error

		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: s.Text(),
		})

		resp, err := client.CreateChatCompletion(ctx, req)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)

		fmt.Println(resp.Choices[0].FinishReason)

		if resp.Choices[0].FinishReason == openai.FinishReasonToolCalls {
			fmt.Println("Calling function!")
			name := fmt.Sprint(resp.Choices[0].Message.ToolCalls[0].Function.Name)

			var args AiFunc
			err := json.Unmarshal([]byte(resp.Choices[0].Message.ToolCalls[0].Function.Arguments), &args)

			if err != nil {
				fmt.Printf("Issue with args: %v\n", err)
			}
			// call func
			fmt.Printf("calling func: %s\n", name)
			ff.endGame(ctx, client, args.Prompt)
			os.Exit(0)
		}
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
