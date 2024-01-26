package services

import (
	"context"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
	"os"
)

func GeminiTextGenFunction(text string) *genai.Content {
	apiKey := os.Getenv("GEMINI_API_KEY")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text(text))
	if err != nil {
		log.Fatal(err)
	}

	for _, candidate := range resp.Candidates {
		if candidate != nil {
			return candidate.Content
		}
	}
	return nil
}

func GeminiImageToTextFunction(imgData []byte, text string) *genai.Content {

	apiKey := os.Getenv("GEMINI_API_KEY")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro-vision")

	prompt := []genai.Part{
		genai.ImageData("jpeg", imgData),
		genai.Text(text),
	}
	resp, err := model.GenerateContent(ctx, prompt...)

	for _, candidate := range resp.Candidates {
		if candidate != nil {
			return candidate.Content
		}
	}
	return nil

}
