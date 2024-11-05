package version

import (
	"fmt"
	"sii/pkg/build"

	"github.com/spf13/cobra"
)

func VersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Current version of the SII application.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s\n", build.VERSION)
			fmt.Printf("Build:   %s\n", build.BUILD)
			fmt.Printf("Commit:  %s\n", build.COMMIT)
		},
	}
}
