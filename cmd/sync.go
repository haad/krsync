package cmd

import (
	"fmt"
	// "os"
	"time"

	"github.com/spf13/cobra"
	"github.com/zloylos/grsync"
)

func init() {
	var syncCmd = &cobra.Command{
		Use:   "sync",
		Short: "Sync given directory to a container running inside Kubernetes.",
		Run: func(cmd *cobra.Command, args []string) {
			syncFiles()
		},
	}
	rootCmd.AddCommand(syncCmd)

	// var projCreateCmd = &cobra.Command{
	//     Use:   "create",
	//     Short: "Create project with given name",
	//     Long:  `Create project which belongs to one customer`,
	//     Run: func(cmd *cobra.Command, args []string) {
	//         project.ProjectCreate(projName, projEstimate, projCustName)
	//     },
	// }
	// projCreateCmd.Flags().StringVarP(&projName, "name", "n", "", "Project name")
	// projCreateCmd.Flags().StringVarP(&projCustName, "customer", "c", "", "Project customer name, needs to be created before.")
	// projCreateCmd.Flags().StringVarP(&projEstimate, "estimate", "e", "", "Project hour estimate, valid units are s/m/h")
	// projCreateCmd.MarkFlagRequired("name")
	// projCreateCmd.MarkFlagRequired("customer")
	// projCmd.AddCommand(projCreateCmd)

}

func syncFiles() {
	task := grsync.NewTask(
		"username@server.com:/source/folder",
		"/home/user/destination",
		grsync.RsyncOptions{},
	)

	go func() {
		for {
			state := task.State()
			fmt.Printf(
				"progress: %.2f / rem. %d / tot. %d / sp. %s \n",
				state.Progress,
				state.Remain,
				state.Total,
				state.Speed,
			)
			<-time.After(time.Second)
		}
	}()

	if err := task.Run(); err != nil {
		panic(err)
	}

	fmt.Println("well done")
	fmt.Println(task.Log())
}
