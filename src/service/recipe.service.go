package service

import (
	"cloud.google.com/go/firestore"
	"github.com/gdsc-ys/lookncook-server/internal/utils"
	"log"
)

type RecipeService interface {
	GenerateRecipeFromFridgeContent(imgData []byte) (string, string)
}

type RecipeServiceImpl struct {
	geminiService   GeminiService
	firestoreClient *firestore.Client
}

func (service *RecipeServiceImpl) GenerateRecipeFromFridgeContent(imgData []byte) (string, string) {
	// fridge Content
	fridgeContent := service.generateFridgeContent(imgData)

	// Use the fridge content as the prompt for generating the recipe
	promptText, _ := utils.ReadFileAsString("internal/prompt/recipePrompt.txt")
	recipePrompt := fridgeContent + promptText
	recipeContent := service.generateRecipe(recipePrompt)

	return fridgeContent, recipeContent
}

func (service *RecipeServiceImpl) generateFridgeContent(imgData []byte) string {
	promptText, err := utils.ReadFileAsString("internal/prompt/fridgeContentPrompt.txt")
	if err != nil {
		log.Fatal("Error occurred while reading prompt : " + err.Error())
	}

	content := service.geminiService.GenerateTextFromImage(imgData, promptText)
	fridgeContent := utils.GeminiContentToString(content)

	return fridgeContent
}

func (service *RecipeServiceImpl) generateRecipe(fridgeContent string) string {

	recipe := service.geminiService.GenerateTextFromText(fridgeContent)
	recipeContent := utils.GeminiContentToString(recipe)

	return recipeContent
}

func RecipeServiceInit(geminiService GeminiService, firestoreClient *firestore.Client) *RecipeServiceImpl {
	return &RecipeServiceImpl{
		geminiService:   geminiService,
		firestoreClient: firestoreClient,
	}
}
