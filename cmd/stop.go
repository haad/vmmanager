package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	var vmx string
	var stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop virtual machine ",
		Long:  "This command stops vmware fusion virtual machine.",
		Run: func(cmd *cobra.Command, args []string) {
			// Execute the "vmrun" command in the current directory
			output, err := exec.Command("vmrun", "stop", vmx).Output()
			if err != nil {
				fmt.Println("Error Message:", string(output))
				fmt.Println("Error executing command:", err)
				os.Exit(1)
			}

			// Print the command output
			fmt.Println(string(output))
		},
	}

	stopCmd.Flags().StringVarP(&vmx, "vmx", "V", "", "VM vmx file")
	stopCmd.MarkFlagRequired("vmx")
	rootCmd.AddCommand(stopCmd)
}
