package resource

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"os"
)

const (
	TableResourceURI = "docs://table"
)

// 提供表结构
func TableResource() mcp.Resource {
	return mcp.NewResource(
		TableResourceURI,
		"mysql table info",
		mcp.WithResourceDescription("本项目的mysql表结构"),
		mcp.WithMIMEType("text/plain"),
		mcp.ResourceContents(),
	)
}

func TableResourceFunc(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	// 读取表结构
	sqlBytes, err := os.ReadFile("tech.sql")
	if err != nil {
		return nil, fmt.Errorf("failed to read sql file: %w", err)
	}
	sqlContent := string(sqlBytes)

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      TableResourceURI,
			MIMEType: "text/plain",
			Text:     sqlContent,
		},
	}, nil
}
