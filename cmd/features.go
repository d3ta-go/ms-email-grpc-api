package cmd

import (
	"github.com/d3ta-go/ms-email-grpc-api/cmd/db"
	"github.com/d3ta-go/ms-email-grpc-api/cmd/server"
)

func init() {
	RootCmd.AddCommand(db.DBCmd)
	RootCmd.AddCommand(server.ServerCmd)
	// Add your custom command
}
