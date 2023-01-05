package cmd

import (
	"github.com/haad/vmmanager/vmware"
	"github.com/spf13/cobra"
)

func init() {
	var unpauseCmd = &cobra.Command{
		Use:   "unpause [path_to_file.vmx]",
		Short: "Unpause virtual machine ",
		Long:  "This command unpauses vmware fusion virtual machine.",
		Args: func(cmd *cobra.Command, args []string) error {
			return validatePostArguments(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			f := vmware.NewFusion(args[0], "", "", "")
			f.Unpause()
		},
	}

	rootCmd.AddCommand(unpauseCmd)
}
