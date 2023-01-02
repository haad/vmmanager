package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	var gui bool
	var startCmd = &cobra.Command{
		Use:   "start [--gui] [path_to_file.vmx]",
		Short: "Start virtual machine ",
		Long:  "This command starts vmware fusion virtual machine.",
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
			var vmFlags = vmrunFlags{Hard: false, Soft: false, Gui: gui}

			vmrunExecCommand("start", args[0], &vmFlags)
		},
	}

	startCmd.PersistentFlags().BoolVar(&gui, "gui", false, "Enables gui window for vm start.")
	//startCmd.Flags().StringVarP(&vmx, "vmx", "V", "", "VM vmx file")
	//startCmd.MarkFlagRequired("vmx")
	rootCmd.AddCommand(startCmd)

}
