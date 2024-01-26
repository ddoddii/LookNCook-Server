package services

import (
	"lookncook/internal/utils"
)

func recipeGenFunction(fridgeContent string) string {

	recipe := GeminiTextGenFunction(fridgeContent)
	var recipeContent string = utils.GeminiContentToString(recipe)

	return recipeContent
}

func GenerateRecipeFromFridgeContent(imgData []byte) (string, string) {
	// fridge Content
	fridgeContent := FridgeContentGenFunction(imgData)

	// Use the fridge content as the prompt for generating the recipe
	promptText, _ := utils.ReadFileAsString("../../internal/prompt/recipePrompt.txt")
	recipeContent := recipeGenFunction(fridgeContent + promptText)

	return fridgeContent, recipeContent
}
