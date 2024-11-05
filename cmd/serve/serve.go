package serve

import (
	"github.com/spf13/cobra"
	"log"
	"sii/pkg/serve"
)

func ServeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the queue server.",
		Run: func(cmd *cobra.Command, args []string) {
			server := serve.Server{}

			err := server.Start(cmd.Context())
			if err != nil {
				log.Printf("server quit: %v", err)
			}
			log.Println("server quit without error")
		},
	}
}
