package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	var hardR bool
	var softR bool
	var suspendCmd = &cobra.Command{
		Use:   "suspend [--hard] [path_to_file.vmx]",
		Short: "Suspends virtual machine ",
		Long:  "This command suspends vmware fusion virtual machine.",
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

			vmrunExecCommand("suspend", args[0], &vmFlags)
		},
	}

	// suspendCmd.Flags().StringVarP(&vmx, "vmx", "V", "", "VM vmx file")
	// suspendCmd.MarkFlagRequired("vmx")
	suspendCmd.PersistentFlags().BoolVar(&hardR, "hard", false, "Forces VM to do a hard suspend.")
	suspendCmd.PersistentFlags().BoolVar(&softR, "soft", true, "Does soft stop for a vm with signal to shutdown.")
	rootCmd.AddCommand(suspendCmd)
}
