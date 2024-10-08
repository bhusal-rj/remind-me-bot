package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/bhusal-rj/remind-me/config"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GetInfoFromGemini() string {
	ctx := context.Background()

	gemini_key := config.InitialConfig.Gemini_Key
	client, err := genai.NewClient(ctx, option.WithAPIKey(gemini_key))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	prompt := []genai.Part{
		genai.Text(config.InitialConfig.PROMPT),
	}
	response, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// responseFromGemini := generateResponse(response)
	resp := response.Candidates[0].Content.Parts[0]
	respJson, err := json.Marshal(resp)
	return string(respJson)
}
