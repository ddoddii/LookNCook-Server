package service

import (
	"encoding/json"
	"github.com/gdsc-ys/lookncook-server/internal/utils"
	"github.com/gdsc-ys/lookncook-server/src/model"
	"log"
)

type KitchenService interface {
	GetCookingEnvironmentHazards(imgData []byte) []model.Hazard
}

type KitchenServiceImpl struct {
	geminiService GeminiService
}

func (service *KitchenServiceImpl) GetCookingEnvironmentHazards(imgData []byte) []model.Hazard {
	promptText, err := utils.ReadFileAsString("internal/prompt/CookingEnvironmentPrompt.txt")

	if err != nil {
		log.Fatal("Error occurred while reading prompt : " + err.Error())
	}

	content := service.geminiService.GenerateTextFromImage(imgData, promptText)
	cookingEnvContent := utils.GeminiContentToString(content)

	// Parse Kitchen Env
	var hazards []model.Hazard
	err = json.Unmarshal([]byte(cookingEnvContent), &hazards)
	if err != nil {
		log.Fatalf("Error parsing JSON in Kitchen Environment : %s", err)
	}

	return hazards
}

func KitchenServiceInit(geminiService GeminiService) *KitchenServiceImpl {
	return &KitchenServiceImpl{
		geminiService: geminiService,
	}
}
