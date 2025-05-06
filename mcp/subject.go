package main

import (
	"fmt"
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

	// Prompt
	s.AddPrompt()

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

}
