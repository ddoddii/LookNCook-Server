package services

import (
	"log"
	"lookncook/internal/utils"
)

func FridgeContentGenFunction(imgData []byte) string {
	promptText, err := utils.ReadFileAsString("../../internal/prompt/fridgeContentPrompt.txt")
	if err != nil {
		log.Fatal("Error occurred while reading prompt : " + err.Error())
	}

	content := GeminiImageToTextFunction(imgData, promptText)

	fridgeContent := utils.GeminiContentToString(content)

	return fridgeContent
}
