package prompts

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// 代码检查
func CodeReview() server.PromptHandlerFunc {
	return func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		prNumber := request.Params.Arguments["pr_number"]
		if prNumber == "" {
			return nil, fmt.Errorf("pr_number is required")
		}

		return mcp.NewGetPromptResult(
			"Code review assistance",
			[]mcp.PromptMessage{mcp.NewPromptMessage(
				mcp.RoleUser,
				mcp.NewTextContent("You are a helpful code reviewer. Review the changes and provide constructive feedback."),
			), mcp.NewPromptMessage(
				mcp.RoleAssistant,
				mcp.NewEmbeddedResource(mcp.TextResourceContents{
					URI:      "test://static/resource",
					MIMEType: "text/plain",
					Text:     "This is a sample resource",
				}),
			)},
		), nil

	}
}
