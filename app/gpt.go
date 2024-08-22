package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type ChatRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type ChatCompletionResponse struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type jokeMsg struct {
	joke string
}

func getJoke() tea.Msg {
	url := "https://api.openai.com/v1/chat/completions"
	payload := ChatRequest{
		Model: "gpt-4o",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{Role: "system", Content: "You are a a wisecracking assistant."},
			{Role: "user", Content: "Tell me a joke about software. Please make it different from your previous jokes"},
		},
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return errMsg{err}
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", PluginOptions.Token))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return errMsg{err}
	}

	client := &http.Client{
		Timeout: time.Duration(PluginOptions.Network.Timeout) * time.Millisecond,
	}

	resp, err := client.Do(req)

	if err != nil {
		return errMsg{err}
	}

	if resp.StatusCode != 200 {
		return errMsg{fmt.Errorf("ChatGPT: %s!\n", resp.Status)}
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return errMsg{err}
	}

	var chatResponse ChatCompletionResponse
	err = json.Unmarshal(b, &chatResponse)
	if err != nil {
		return errMsg{err}
	}

	return jokeMsg{chatResponse.Choices[0].Message.Content}
}
