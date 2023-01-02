package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	var gui bool
	var vmx string
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start virtual machine ",
		Long:  "This command starts vmware fusion virtual machine.",
		Run: func(cmd *cobra.Command, args []string) {
			var vmgui string

			if gui {
				vmgui = "gui"
			} else {
				vmgui = "nogui"
			}

			// Execute the "vmrun" command in the current directory
			output, err := exec.Command("vmrun", "start", vmx, vmgui).Output()
			if err != nil {
				fmt.Println("Error Message:", string(output))
				fmt.Println("Error executing command:", err)
				os.Exit(1)
			}

			// Print the command output
			fmt.Println(string(output))
		},
	}

	startCmd.PersistentFlags().BoolVar(&gui, "gui", false, "Enables gui window for vm start")
	startCmd.Flags().StringVarP(&vmx, "vmx", "V", "", "VM vmx file")
	startCmd.MarkFlagRequired("vmx")

	rootCmd.AddCommand(startCmd)

}
