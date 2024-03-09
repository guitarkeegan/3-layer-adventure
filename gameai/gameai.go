package gameai

import (
	"context"
	"errors"

	"github.com/sashabaranov/go-openai"
)

const (
	Chatting = "chatting"
	Ending   = "ending"
)

type GameAI struct {
	state string
}

func New() *GameAI {
	return &GameAI{
		state: Chatting,
	}
}

func (gai *GameAI) State() string {
	return gai.state
}

func (gai *GameAI) Chat(ctx context.Context, c *openai.Client) (string, error) {
	// TODO: check state
	if gai.state == Ending {
		return "", errors.New("Cannot call Chat while in Ending state")

	}
	// TODO: send message to openai
	return "", nil
}
