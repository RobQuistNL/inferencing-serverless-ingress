package cmd

import (
	"os"
	"sii/cmd/serve"
	"sii/cmd/version"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sii",
	Short: "Serverless inferencing ingress server.",
	Long:  `Receive inferencing requests, queue them, and provide them to runners / systems.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd := rootCmd
	commands := append([]*cobra.Command{
		serve.ServeCommand(),
		version.VersionCommand(),
	})
	cmd.AddCommand(commands...)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Persistent flags available throughout the entire application
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sii.yaml)")
}
