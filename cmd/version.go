package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

// Version is a global version string
const Version = "0.0.1"

// GitSHA is latest git HEAD. We want to replace this variable at build time with "-ldflags -X cmd.GitSHA=xxx", where const is not supported.
var GitSHA = ""

func init() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of krsync",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Kubernetes pod Rsync version %s (Git SHA: %s, Go Version: %s)\n", Version, GitSHA, runtime.Version())
		},
	}
	rootCmd.AddCommand(versionCmd)
}
