package claude

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

const anthropicAPI = "https://api.anthropic.com/v1/messages"

type Client struct {
	apiKey string
	http   *http.Client
}

type messageRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type messageResponse struct {
	Content []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"content"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		http:   &http.Client{},
	}
}

func (c *Client) CreateMessage(prompt string) (string, error) {
	// Get the main response
	reqBody := messageRequest{
		Model:     "claude-3-opus-20240229",
		MaxTokens: 1024,
		Messages: []Message{
			{Role: "user", Content: "Please provide a title and a response to the following prompt. The title should be a markdown h1, and the date should be a tag in the format #YYYY-MM-DD. Prompt: " + prompt + "\n\nProvide your response without any '#+Response' prefix. End your response with: 'To modify this note, just let me know how. Otherwise, we can end this session.'"},
		},
	}

	content, err := c.sendRequest(reqBody)
	if err != nil {
		return "", err
	}

	// Clean up response formatting
	content = strings.TrimPrefix(content, "#+Response")
	content = strings.TrimSpace(content)

	return content, nil
}

func (c *Client) sendRequest(reqBody messageRequest) (string, error) {
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", anthropicAPI, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	// Update headers to include both content-type and anthropic-beta
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)
	req.Header.Set("Anthropic-Version", "2023-06-01")
	req.Header.Set("anthropic-beta", "messages-2023-12-15")

	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API request failed with status: %d - %s", resp.StatusCode, string(bodyBytes))
	}

	var result messageResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Content) > 0 {
		// Clean up any URL encoding artifacts and normalize whitespace
		text := result.Content[0].Text
		text = strings.ReplaceAll(text, "+", " ")
		text = strings.ReplaceAll(text, "%20", " ")
		text = strings.ReplaceAll(text, "&amp;", "&") // Decode HTML entities
		text = strings.TrimSpace(text)
		return text, nil
	}

	return "", fmt.Errorf("no content in response")
}

// CreateFromClipboard creates a new Bear note using the current clipboard content
func (c *Client) CreateFromClipboard(title string, tags []string) (string, error) {
	params := make([]string, 0)
	params = append(params, "clipboard=yes")

	if title != "" {
		params = append(params, "title="+strings.ReplaceAll(title, " ", "%20"))
	}

	if len(tags) > 0 {
		encodedTags := strings.Join(tags, ",")
		encodedTags = strings.ReplaceAll(encodedTags, " ", "%20")
		params = append(params, "tags="+encodedTags)
	}

	bearURL := fmt.Sprintf("bear://x-callback-url/create?%s", strings.Join(params, "&"))

	cmd := exec.Command("open", bearURL)
	return "", cmd.Run()
}
