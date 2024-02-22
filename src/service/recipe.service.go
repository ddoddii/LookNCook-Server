package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/gdsc-ys/lookncook-server/internal/utils"
	"github.com/gdsc-ys/lookncook-server/src/constant"
	"github.com/gdsc-ys/lookncook-server/src/model"
	"log"
)

type RecipeService interface {
	GenerateRecipeFromFridgeContent(imgData []byte) ([]model.Ingredient, []model.Recipe)
}

type RecipeServiceImpl struct {
	geminiService   GeminiService
	firestoreClient *firestore.Client
}

func (service *RecipeServiceImpl) GenerateRecipeFromFridgeContent(imgData []byte) ([]model.Ingredient, []model.Recipe) {
	// fridge Content
	fridgeContent, ingredients := service.generateFridgeContent(imgData)
	// Use the fridge content as the prompt for generating the recipe
	promptText, _ := utils.ReadFileAsString("internal/prompt/recipePrompt.txt")
	recipePrompt := fridgeContent + promptText
	_, recipes := service.generateRecipe(recipePrompt)

	return ingredients, recipes
}

func (service *RecipeServiceImpl) generateFridgeContent(imgData []byte) (string, []model.Ingredient) {
	promptText, err := utils.ReadFileAsString("internal/prompt/fridgeContentPrompt.txt")
	if err != nil {
		log.Fatal("Error occurred while reading prompt : " + err.Error())
	}

	content := service.geminiService.GenerateTextFromImage(imgData, promptText)
	fridgeContent := utils.GeminiContentToString(content)

	// Parse Ingredients
	var ingredients []model.Ingredient
	err = json.Unmarshal([]byte(fridgeContent), &ingredients)
	if err != nil {
		log.Fatalf("Error parsing Json in Ingredients : %v", err)
	}

	// Add Ingredients
	service.addIngredients(ingredients)

	return fridgeContent, ingredients
}

func (service *RecipeServiceImpl) generateRecipe(recipePrompt string) (string, []model.Recipe) {

	recipe := service.geminiService.GenerateTextFromText(recipePrompt)
	recipeContent := utils.GeminiContentToString(recipe)

	// Parse Recipe
	var recipes []model.Recipe
	err := json.Unmarshal([]byte(recipeContent), &recipes)
	if err != nil {
		log.Fatalf("Error parsing JSON in Recipes : %s", err)
	}

	// Add Recipe
	service.addRecipes(recipes)

	return recipeContent, recipes
}

// Add ingredients to firestore
func (service *RecipeServiceImpl) addIngredients(ingredients []model.Ingredient) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.FirestoreDefaultTimeout)
	defer cancel()

	for _, ingredient := range ingredients {
		_, _, err := service.firestoreClient.
			Collection(constant.IngredientCollection).
			Add(ctx, ingredient)
		if err != nil {
			log.Printf("Failed to add ingredient: %v", err)
		}
	}

}

// Add recipes to firestore
func (service *RecipeServiceImpl) addRecipes(recipes []model.Recipe) {
	ctx, cancel := context.WithTimeout(context.Background(), constant.FirestoreDefaultTimeout)
	defer cancel()

	for _, recipe := range recipes {
		_, _, err := service.firestoreClient.
			Collection(constant.RecipeCollection).
			Add(ctx, recipe)
		if err != nil {
			log.Printf("Failed to add recipe: %v", err)
		}
	}
}

func RecipeServiceInit(geminiService GeminiService, firestoreClient *firestore.Client) *RecipeServiceImpl {
	return &RecipeServiceImpl{
		geminiService:   geminiService,
		firestoreClient: firestoreClient,
	}
}
