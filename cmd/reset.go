package cmd

import (
	"github.com/haad/vmmanager/vmware"
	"github.com/spf13/cobra"
)

func init() {
	var hardR bool
	var resetCmd = &cobra.Command{
		Use:   "reset [--hard] [path_to_file.vmx]",
		Short: "Reset virtual machine ",
		Long:  "This command resets vmware fusion virtual machine.",
		Args: func(cmd *cobra.Command, args []string) error {
			return validatePostArguments(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			f := vmware.NewFusion(args[0], "", "", "")
			if hardR {
				f.Reset()
			} else {
				f.Restart()
			}
		},
	}

	resetCmd.PersistentFlags().BoolVar(&hardR, "hard", false, "Forces VM to do a hard reset.")
	rootCmd.AddCommand(resetCmd)
}
