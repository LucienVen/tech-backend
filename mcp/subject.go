package main

import (
	"fmt"
	"github.com/LucienVen/tech-backend/mcp/resource"
	"github.com/mark3labs/mcp-go/server"
	"log"
)

func main() {
	log.Println("Starting MCP tech demo...")

	s := server.NewMCPServer(
		"school-teacher-student crm",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	// 添加 mysql 表资源
	s.AddResource(resource.TableResource(), resource.TableResourceFunc)

	// Prompt
	s.AddPrompt()

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

}
