package gameai

import (
	"context"
	"errors"

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

func (gai *GameAI) Chat(ctx context.Context, c *openai.Client) (string, error) {

	if gai.state != Chatting {
		return "", errors.New("Cannot call Chat while in Ending state")

	}
	// TODO: send message to openai

	return "", nil
}

func (gai *GameAI) CallFunction(ctx context.Context, c *openai.Client) {
	// TODO: This will end with state changing to 'ending'
}
