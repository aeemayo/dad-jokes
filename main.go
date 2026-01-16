package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/TeneoProtocolAI/teneo-agent-sdk/pkg/agent"
	"github.com/joho/godotenv"
)

type DadJokeAgent struct{}

type OpenRouterRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenRouterResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

func (a *DadJokeAgent) ProcessTask(ctx context.Context, task string) (string, error) {
	log.Printf("Processing task: %s", task)

	// Clean up the task input
	task = strings.TrimSpace(task)
	task = strings.TrimPrefix(task, "/")
	taskLower := strings.ToLower(task)

	// Split into command and arguments
	parts := strings.Fields(taskLower)
	if len(parts) == 0 {
		return "No command provided. Available commands: humor_me", nil
	}

	command := parts[0]
	// args := parts[1:]

	// Route to appropriate command handler
	switch command {
	case "humor_me":
		joke, err := getDadJokeFromOpenRouter()
		if err != nil {
			log.Printf("Error fetching joke: %v", err)
			return "I tried to think of a joke, but I got an error instead. Why did the server break up with the network? Because it couldn't find a connection.", nil
		}
		return joke, nil

	default:
		// NLP Fallback enabled: This input didn't match any commands
		// TODO: Implement NLP processing for natural language input
		// Users can send natural language queries like: @dad-joke "what's the weather?"
		//
		// Implementation options:
		// - Integrate with OpenAI API (see OpenAI agent example)
		// - Use your own ML model
		// - Call external NLP services
		// - Implement custom text processing logic
		//
		// For now, returning placeholder response:
		return fmt.Sprintf("NLP processing not yet implemented. Received: %s", task), nil
	}
}

func getDadJokeFromOpenRouter() (string, error) {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENROUTER_API_KEY is not set")
	}

	reqBody := OpenRouterRequest{
		Model: "openai/gpt-3.5-turbo",
		Messages: []Message{
			{Role: "system", Content: "You are a dad joke expert. Tell me a short, classic dad joke."},
			{Role: "user", Content: "Tell me a dad joke."},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("HTTP-Referer", "https://github.com/TeneoProtocolAI/teneo-agent-sdk")
	req.Header.Set("X-Title", "Dad Joke Agent")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var openRouterResp OpenRouterResponse
	if err := json.NewDecoder(resp.Body).Decode(&openRouterResp); err != nil {
		return "", err
	}

	if len(openRouterResp.Choices) > 0 {
		return openRouterResp.Choices[0].Message.Content, nil
	}

	return "No joke found.", nil
}

func main() {
	godotenv.Load()
	config := agent.DefaultConfig()

	config.Name = "dad jokes"
	config.Description = `A great beginner-friendly agent is a Dad Joke Teller. Users send any 
message (e.g., "Tell me a joke about computers" or just "Hi"), and the agent responds with a 
groon-worthy, family-friendly dad joke — ideally related to the topic, or a random one if not.`
	config.Capabilities = []string{"humor" , "comedy"}
	config.PrivateKey = os.Getenv("PRIVATE_KEY")
	config.NFTTokenID = os.Getenv("NFT_TOKEN_ID")
	config.OwnerAddress = os.Getenv("OWNER_ADDRESS")

	enhancedAgent, err := agent.NewEnhancedAgent(&agent.EnhancedAgentConfig{
		Config:       config,
		AgentHandler: &DadJokeAgent{},
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting dad jokes...")
	enhancedAgent.Run()
}