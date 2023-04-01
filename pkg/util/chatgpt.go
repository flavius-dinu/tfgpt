package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const ChatGPTAPIURL = "https://api.openai.com/v1/chat/completions"

func GetExplanationFromChatGPT(output string, prompt_builder string, command string) (string, error) {

	apiKey, err := GetAPIKey()
	if err != nil {
		return "", err
	}
	var requestBody []byte

	if prompt_builder == "command" {
		prompt := fmt.Sprintf("I've ran this terraform command %s. Explain the following Terraform command output. If there are errors encountered, please provide remediation otherwise omit this instruction. This is the command output:\n\n%s", command, output)

		requestBody, err = json.Marshal(map[string]interface{}{
			"model": "gpt-3.5-turbo",
			"messages": []map[string]string{
				{
					"role":    "system",
					"content": "You are an expert Terraform engineer that explains what happens when a Terraform command runs.",
				},
				{
					"role":    "user",
					"content": prompt,
				},
			},
			"max_tokens": 500,
		})
		if err != nil {
			return "", err
		}
	} else if prompt_builder == "concept" {
		prompt := fmt.Sprintf("Can you explain what this terraform concept is? If it is not used in Terraform, please mention my confusion.:\n\n%s", output)

		requestBody, err = json.Marshal(map[string]interface{}{
			"model": "gpt-3.5-turbo",
			"messages": []map[string]string{
				{
					"role":    "system",
					"content": "You are an expert Terraform engineer that explains Terraform concepts.",
				},
				{
					"role":    "user",
					"content": prompt,
				},
			},
			"max_tokens": 500,
		})
		if err != nil {
			return "", err
		}
	} else {
		return "", fmt.Errorf("invalid prompt_builder value: %s", prompt_builder)
	}

	req, err := http.NewRequest("POST", ChatGPTAPIURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responseObj map[string]interface{}

	if err := json.Unmarshal(responseBody, &responseObj); err != nil {
		return "", err
	}

	if errMsg, ok := responseObj["error"].(map[string]interface{}); ok {
		if msg, ok := errMsg["message"].(string); ok {
			return "", fmt.Errorf("ChatGPT API error: %s", msg)
		}
	}

	if choices, ok := responseObj["choices"].([]interface{}); ok && len(choices) > 0 {
		if choice, ok := choices[0].(map[string]interface{}); ok {
			if message, ok := choice["message"].(map[string]interface{}); ok {
				if text, ok := message["content"].(string); ok {
					return strings.TrimSpace(text), nil
				}
			}
		}
	}

	return "", fmt.Errorf("failed to parse ChatGPT response")
}

func GetAPIKey() (string, error) {
	// Try to get the API key from the environment variable
	apiKey := os.Getenv("CHATGPT_API_KEY")
	if apiKey != "" {
		return apiKey, nil
	}

	// If not found, try to read the API key from the file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %s", err)
	}

	credentialsPath := filepath.Join(homeDir, ".tfgpt", "credentials")
	apiKey, err = ReadAPIKeyFromFile(credentialsPath)
	if err != nil {
		return "", fmt.Errorf("failed to read API key from file: %s", err)
	}

	return apiKey, nil
}

func ReadAPIKeyFromFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	apiKey := strings.TrimSpace(string(content))
	if apiKey == "" {
		return "", fmt.Errorf("API key is empty")
	}

	return apiKey, nil
}
