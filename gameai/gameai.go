package gameai

import (
	"3-layer-adventure/gameai/functions"
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

const (
	Loading  GameState = "loading"
	Chatting GameState = "chatting"
	Ending   GameState = "ending"
)

type GameState string

type GameAI struct {
	state GameState
}

func New() *GameAI {
	return &GameAI{
		state: Chatting,
	}
}

func (gai *GameAI) State() GameState {
	return gai.state
}

func (gai *GameAI) Chat(ctx context.Context, c *openai.Client, s *bufio.Scanner, req *openai.ChatCompletionRequest) (string, error) {

	if gai.state != Chatting {
		return "", errors.New("Cannot call Chat while in Ending state")

	}

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

		resp, err := c.CreateChatCompletion(ctx, *req)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)

		fmt.Println(resp.Choices[0].FinishReason)
		// if state == ending then call the function

		if resp.Choices[0].FinishReason == openai.FinishReasonToolCalls {
			gai.state = Ending
		}

		if gai.state == Ending {
			fmt.Println("Calling function!")
			// ToolCalls[0] related to my Tools?
			name := fmt.Sprint(resp.Choices[0].Message.ToolCalls[0].Function.Name)

			var args functions.AiFunc
			err := json.Unmarshal([]byte(resp.Choices[0].Message.ToolCalls[0].Function.Arguments), &args)

			if err != nil {
				fmt.Printf("Issue with args: %v\n", err)
			}
			// call func
			fmt.Printf("calling func: %s\n", name)
			gai.CallFunction(ctx, c, name, args.Prompt)
			break
		}

		req.Messages = append(req.Messages, resp.Choices[0].Message)
		fmt.Print("> ")
	}

	return "", nil
}

func (gai *GameAI) CallFunction(ctx context.Context, c *openai.Client, n string, p string) {

	ff := new(functions.FunkFactory)
	ff.EndGame(ctx, c, p)
	os.Exit(0)

}
