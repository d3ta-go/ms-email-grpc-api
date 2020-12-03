package server

import (
	"fmt"

	"github.com/d3ta-go/ms-email-grpc-api/interface/http-apps/grpc/micro"
	"github.com/spf13/cobra"
)

// grpcAPICmd represents the grpc API server command
var grpcAPICmd = &cobra.Command{
	Use:   "grpcapi",
	Short: "Shows the grpcapi server command.",
	Long:  `Shows the grpcapi server command.`,
	Run: func(cmd *cobra.Command, args []string) {

		if err := micro.StartGrpcAPIServer(); err != nil {
			fmt.Println("Error while running `StartGrpcAPIServer()`: ")
			panic(err)
		}
	},
}

func init() {
	ServerCmd.AddCommand(grpcAPICmd)
}
