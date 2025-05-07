package prompts

import "github.com/mark3labs/mcp-go/mcp"

func NewPrompt() []mcp.Prompt {

}

func QueryBuilderPrompt() mcp.Prompt {
	return mcp.NewPrompt(
		"query_builder",
		mcp.WithPromptDescription("SQL query builder assistance"),
		mcp.WithArgument(
			"table",
			mcp.ArgumentDescription("Name of the table to query"),
			mcp.RequiredArgument(),
		),
	)
}
