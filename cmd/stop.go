package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	var hardR bool
	var softR bool
	var stopCmd = &cobra.Command{
		Use:   "stop [--hard] [path_to_file.vmx]",
		Short: "Stop virtual machine ",
		Long:  "This command stops vmware fusion virtual machine.",
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
			var vmFlags = vmrunFlags{Hard: hardR, Soft: softR, Gui: false}

			vmrunExecCommand("stop", args[0], &vmFlags)
		},
	}

	// stopCmd.Flags().StringVarP(&vmx, "vmx", "V", "", "VM vmx file")
	// stopCmd.MarkFlagRequired("vmx")
	stopCmd.PersistentFlags().BoolVar(&hardR, "hard", false, "Forces VM to do a hard shutdown.")
	stopCmd.PersistentFlags().BoolVar(&softR, "soft", true, "Does soft stop for a vm with signal to shutdown.")
	rootCmd.AddCommand(stopCmd)
}
