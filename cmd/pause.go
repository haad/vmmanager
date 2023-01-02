package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	var pauseCmd = &cobra.Command{
		Use:   "pause [path_to_file.vmx]",
		Short: "Pause virtual machine ",
		Long:  "This command pauses vmware fusion virtual machine.",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}

			if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
				return err
			}

			parts := strings.Split(args[0], ".")
			if strings.Contains(parts[len(parts)-1], "vmx") {
				return nil
			}

			return fmt.Errorf("invalid vmx path specified: %s", args[0])
		},
		Run: func(cmd *cobra.Command, args []string) {
			vmrunExecCommand("pause", args[0], nil)
		},
	}

	rootCmd.AddCommand(pauseCmd)
}
