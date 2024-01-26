package services

import "github.com/google/generative-ai-go/genai"

func recipeGenFunction(fridgeContent string) *genai.Content {

	recipe := GeminiTextGenFunction(fridgeContent)
	return recipe
}

func GenerateRecipeFromFridgeContent(imgData []byte) (*genai.Content, error) {

	fridgeContent := FridgeContentGenFunction(imgData)

	// Use the fridge content as the prompt for generating the recipe
	recipeContent := recipeGenFunction(fridgeContent)

	return recipeContent, nil
}
