package utils

import (
	"encoding/json"
	"github.com/google/generative-ai-go/genai"
	"strings"
)

func GeminiContentToString(content *genai.Content) string {
	var stringBuilder strings.Builder

	for _, part := range content.Parts {
		partJSON, err := json.Marshal(part)
		if err != nil {
			// Handle the error appropriately
			continue // or log the error, or return an error from the function
		}

		cleanedPart := cleanResponse(string(partJSON))
		stringBuilder.WriteString(cleanedPart)
		stringBuilder.WriteString(" ")
	}

	return stringBuilder.String()
}

func cleanResponse(resp string) string {
	resp = strings.Replace(resp, "\\n", "\n", -1)
	resp = strings.Replace(resp, "\\\"", "\"", -1)
	resp = strings.Replace(resp, "\\", "", -1)

	return resp
}
