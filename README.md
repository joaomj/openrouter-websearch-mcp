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
| `MODEL_NAME`         | OpenRouter model name  | No       | `google/gemini-2.5-pro-preview-03-25`         |

## Installation 

To build and install the MCP server:

```bash
cd /path/to/openrouter-websearch-mcp
go build -o /path/to/output/openrouter-mcp-server
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

## Cline Integration Setup

1. Build the server:
```bash
cd /path/to/openrouter-websearch-mcp
go build -o /path/to/output/openrouter-mcp-server
```

2. Configure Cline's MCP settings by editing:
`path/to/cline_mcp_settings.json`

3. Add this configuration:
```json
{
  "mcpServers": {
    "openrouter-mcp-server": {
      "command": "/path/to/openrouter-mcp-server",
      "args": [],
      "env": {
        "OPENROUTER_API_KEY": "your-api-key-here"
      },
      "disabled": false,
      "autoApprove": []
    }
  }
}
```

4. Restart Cline extension to load the new MCP server

5. Test the connection by using the `search_web` tool in Cline
