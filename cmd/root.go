package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vmmanager",
	Short: "Vmware Fusion virtual machine manager",
	Long:  "Simple applicaiton used to download and create new vmware fusion machines. User can also use ti to start/stop/pause and ssh to it.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute runs parses cobra commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	var debug bool
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enables debug logging")
	rootCmd.PersistentFlags().Int("port", 8080, "The port to listen on")
	rootCmd.PersistentFlags().String("username", "", "The username to use")
}
