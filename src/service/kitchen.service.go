package service

import (
	"github.com/gdsc-ys/lookncook-server/internal/utils"
	"log"
)

type KitchenService interface {
	GetCookingEnvironment(imgData []byte) (string, error)
}

type KitchenServiceImpl struct {
	geminiService GeminiService
}

func (service *KitchenServiceImpl) GetCookingEnvironment(imgData []byte) (string, error) {
	promptText, err := utils.ReadFileAsString("internal/prompt/CookingEnvironmentPrompt.txt")

	if err != nil {
		log.Fatal("Error occurred while reading prompt : " + err.Error())
	}

	content := service.geminiService.GenerateTextFromImage(imgData, promptText)
	cookingEnvContent := utils.GeminiContentToString(content)

	return cookingEnvContent, nil
}

func KitchenServiceInit(geminiService GeminiService) *KitchenServiceImpl {
	return &KitchenServiceImpl{
		geminiService: geminiService,
	}
}
