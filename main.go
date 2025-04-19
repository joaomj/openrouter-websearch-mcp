package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ChristianSch/openrouter-websearch-mcp/openrouter"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const prompt = `You are a web search assistant. Based on the following query, provide a concise, well-informed answer using up-to-date knowledge from online sources.

Query: %s

Only include relevant facts that are recent and trustworthy. If something is speculative or unclear, say so.

Answer:
`

func searchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Get the query parameter
	query, ok := request.Params.Arguments["query"].(string)
	if !ok {
		return nil, errors.New("query must be a string")
	}

	// Access API key and model from environment
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	modelName := os.Getenv("MODEL_NAME")
	if apiKey == "" {
		return nil, fmt.Errorf("missing OPENROUTER_API_KEY environment variable")
	}

	if modelName == "" {
		modelName = "google/gemini-2.5-pro-preview-03-25"
	}

	println("Using model:", modelName)

	// Prepare the prompt with the user query
	fullPrompt := fmt.Sprintf(prompt, map[string]interface{}{
		"user_query": query,
	})

	answer, err := openrouter.CallOpenRouter(apiKey, modelName, fullPrompt)
	if err != nil {
		return nil, err
	}

	fmt.Println("Response:", answer)

	return mcp.NewToolResultText(answer), nil
}

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"Openrouter Websearch MCP",
		"0.1.0",
	)

	// Add tool
	tool := mcp.NewTool("search_web",
		mcp.WithDescription("Perform a web search using OpenRouter with Gemini model"),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("The search query to answer"),
		),
	)

	// Add tool handler
	s.AddTool(tool, searchHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
