package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const openAIAPIKey = "YOUR_OPENAI_API_KEY"

type OpenAIRequest struct {
	Model    string   `json:"model"`
	Messages []string `json:"messages"`
}

func GeneratePayload(vulnType, target string) string {
	data := OpenAIRequest{
		Model: "gpt-4",
		Messages: []string{
			fmt.Sprintf("Generate a highly effective %s payload for testing %s.", vulnType, target),
		},
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openAIAPIKey)

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
