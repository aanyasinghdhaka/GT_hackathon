package util

import (
	"context"
	"encoding/base64"
	"os"

	openai "github.com/openai/openai-go"
)

var dalleClient = openai.NewClient(os.Getenv("OPENAI_API_KEY"))

func GenerateBackgroundImages(prompt string) ([][]byte, error) {

	resp, err := dalleClient.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt: prompt,
			N:      4, // generate 4 variants
			Size:   "1024x1024",
			Model:  openai.CreateImageModelDallE3,
		},
	)

	if err != nil {
		return nil, err
	}

	var images [][]byte
	for _, data := range resp.Data {
		b, _ := base64.StdEncoding.DecodeString(data.B64JSON)
		images = append(images, b)
	}

	return images, nil
}
