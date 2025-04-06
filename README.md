# OpenRouter Web Search MCP Server

This project implements an [MCP (Model Context Protocol)](https://github.com/mark3labs/mcp-go) server that provides a `search_web` tool. The tool performs a web search using OpenRouter's online models to generate concise, up-to-date answers.

## Features

- Exposes an MCP-compatible server over stdio
- Provides a `search_web` tool that:
  - Accepts a search query string
  - Uses OpenRouter's online models to generate a fact-based answer
  - Returns the answer as plain text

## Requirements

- Go 1.18+
- An OpenRouter API key

## Environment Variables

| Variable             | Description                                               | Required | Default                                               |
|----------------------|-----------------------------------------------------------|----------|-------------------------------------------------------|
| `OPENROUTER_API_KEY`| Your OpenRouter API key                                   | **Yes**  |                                                       |
| `MODEL_NAME`         | OpenRouter model name (should support `:online` suffix)  | No       | `google/gemini-2.5-pro-exp-03-25:free:online`         |

## Installation 

To install the MCP server, you can use the following command:

```bash
go install github.com/ChristianSch/openrouter-websearch-mcp
```

## Usage

Run the server:

```bash
{
    "mcpServers": {
        "web-search": {
            "command": "openrouter-websearch-mcp",
            "env": {
                "OPENROUTER_API_KEY": "sk-or-v1-..."
            }
        }
    }
}
```