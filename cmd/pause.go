package cmd

import (
	"github.com/haad/vmmanager/vmware"
	"github.com/spf13/cobra"
)

func init() {
	var pauseCmd = &cobra.Command{
		Use:   "pause [path_to_file.vmx]",
		Short: "Pause virtual machine ",
		Long:  "This command pauses vmware fusion virtual machine.",
		Args: func(cmd *cobra.Command, args []string) error {
			return validatePostArguments(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			f := vmware.NewFusion(args[0], "", "", "")
			f.Pause()
		},
	}

	rootCmd.AddCommand(pauseCmd)
}
