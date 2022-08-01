package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

  log "github.com/haad/krsync/log"
)

var rootCmd = &cobra.Command{
	Use:   "krsync",
	Short: "Kubectl Rsync wrapper, uploading files to a container image runing in kubernetes.",
	Long:  `Fast way how to upload files to running container in Kubernetes. `,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute Main command execution entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var debug bool
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enables debug logging")

  log.InitLogger(debug)
}
