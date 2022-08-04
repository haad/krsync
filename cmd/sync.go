package cmd

import (
	"fmt"
	// "os"
	"time"
    //"strings"

	"github.com/spf13/cobra"
	"github.com/zloylos/grsync"

    log "github.com/haad/krsync/log"
)

// /opt/homebrew/bin/rsync -av --progress --stats -vvv --rsh "dist/helper.sh da-dev php da-dev-v33717-da-86b99cc46d-92qbd" test/dist rsync:/tmp/t



// p_user=$USER;p_name=$POD_NAME; rsync -avurP --blocking-io --rsync-path= --rsh="$(which kubectl) exec $p_name -i -- " /home/$USER/target_dir rsync:/home/$p_user/
// krsync -av --progress --stats src-dir/ pod@namespace:/dest-dir
// krsync -av --progress --stats src-dir/ pod:/dest-dir
// /opt/homebrew/bin/rsync -avvv --blocking-io --rsync-path= --rsh="kubectl exec da-dev-v39721-da-c74679766-hcrrc -i -n da-dev -c php --"  test/dist rsync:/tmp/
func init() {
    var namespace string
    var pod string
    var container string
    var progress bool

	var syncCmd = &cobra.Command{
		Use:   "sync",
		Short: "Sync given directory to a container running inside Kubernetes.",
		Run: func(cmd *cobra.Command, args []string) {
            if (len(args) < 2){
                log.Slog.Fatalf("We need src and dst directory paths..")
            }
            src := args[0]
            dst := args[1]
			syncFiles(namespace, pod, container, src, dst, progress)
		},
	}

    syncCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Namespace pod is running in")
    syncCmd.Flags().StringVarP(&pod, "pod", "p", "", "Name of pod where we want to sync our data")
    syncCmd.Flags().StringVarP(&container, "container", "c", "", "Specific container inside given pod if needed.")
    syncCmd.Flags().BoolVarP(&progress, "progress", "P", false, "Print upload files progress.")


    // syncCmd.Flags().StringVarP(&srcdir, "src", "s", "", "Name of pod where we want to sync our data")
    // syncCmd.Flags().StringVarP(&dstdir, "dst", "d", "", "Specific container inside given pod if needed.")


    // syncCmd.MarkFlagRequired("name")
    syncCmd.MarkFlagRequired("pod") // #nosec

	rootCmd.AddCommand(syncCmd)

}

func syncFiles(namespace string, pod string, container string, srcdir string, dstdir string, progress bool) {
    log.Slog.Infof("Uploading files from srcdir: %s, to dstdir: %s, on pod: %s, container: %s in namespace: %s.", srcdir, dstdir, pod, container, namespace)

    rsh := fmt.Sprintf("kubectl exec -i -n %s -c %s %s --", namespace, container, pod)
	task := grsync.NewTask(
		srcdir,
		dstdir,
		grsync.RsyncOptions{
            Verbose: true,
            Archive: true,
            Stats: true,
            Progress: true,
            Compress: true,
            Update: true,
            Recursive: true,
            BlockingIO: true,
            Rsh: rsh,
            RsyncPath: "=", // We have to pass --rsync-path= argument to make sure kubectl exec only has one rsync.
            Exclude: []string {".git"},
        },
	)

    log.Slog.Infof("Rsync cmd: %s", task.GetCmdString() )

    if (progress) {
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
    }

	if err := task.Run(); err != nil {
		panic(err)
	}

	fmt.Println("well done")
	fmt.Println(task.Log())
}
