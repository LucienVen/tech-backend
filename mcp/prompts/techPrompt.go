package prompts

import "github.com/mark3labs/mcp-go/mcp"

func QueryBuilderPrompt() mcp.Prompt {
	return mcp.NewPrompt(
		"query_builder",
		mcp.WithPromptDescription("SQL query builder assistance"),
		mcp.WithArgument(
			"mysql",
			mcp.ArgumentDescription("MySQL table structure"),
			mcp.ArgumentDescription("a"),
			mcp.RequiredArgument(),
		),
	)
}
