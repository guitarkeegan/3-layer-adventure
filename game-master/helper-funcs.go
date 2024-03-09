package gamemaster

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type Funker interface {
	endGame(ctx context.Context, client *openai.Client, prompt string)
}

type AiFunc struct {
	Prompt string `json:"prompt"`
}

type FunkFactory struct {
	funks map[string]Funker
}

func (ff FunkFactory) endGame(ctx context.Context, client *openai.Client, prompt string) {
	fmt.Println("Inside endGame...")
	respUrl, err := client.CreateImage(
		ctx,
		openai.ImageRequest{
			Prompt: prompt, Size: openai.CreateImageSize1024x1024,
			ResponseFormat: openai.CreateImageResponseFormatURL,
			N:              1,
		},
	)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return
	}
	fmt.Println(respUrl.Data[0].URL)
}
