package gamemaster

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func win(client *openai.Client, prompt string) {

	respUrl, err := client.CreateImage(
		context.Background(),
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
