package utils

import (
	"encoding/json"
	"github.com/google/generative-ai-go/genai"
	"strings"
)

func GeminiContentToString(content *genai.Content) string {
	var stringBuilder strings.Builder

	for _, part := range content.Parts {
		// Convert the part to JSON
		partJSON, err := json.Marshal(part)
		if err != nil {
			// Handle the error appropriately
			continue // or log the error, or return an error from the function
		}

		// Append the JSON string
		stringBuilder.Write(partJSON)
		stringBuilder.WriteString(" ") // Adding space between parts
	}

	return stringBuilder.String()
}
