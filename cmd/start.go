package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	var startCmd = &cobra.Command{
		Use:   "exec",
		Short: "Execute a shell command",
		Long:  "This command executes a shell command and prints the output to the console.",
		Run: func(cmd *cobra.Command, args []string) {
			// Execute the "ls" command in the current directory
			output, err := exec.Command("ls").Output()
			if err != nil {
				fmt.Println("Error executing command:", err)
				os.Exit(1)
			}

			// Print the command output
			fmt.Println(string(output))
		},
	}

	rootCmd.AddCommand(startCmd)

}
