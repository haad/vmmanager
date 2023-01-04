package cmd

import (
	"github.com/haad/vmmanager/util"
	"github.com/spf13/cobra"
)

func init() {
	var hardR bool
	var softR bool
	var resetCmd = &cobra.Command{
		Use:   "reset [--hard] [path_to_file.vmx]",
		Short: "Reset virtual machine ",
		Long:  "This command resets vmware fusion virtual machine.",
		Args: func(cmd *cobra.Command, args []string) error {
			return validatePostArguments(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			var vmFlags = util.VmrunFlags{Hard: hardR, Soft: softR, Gui: false}

			util.VmrunExecCommand("reset", args[0], &vmFlags)
		},
	}

	resetCmd.PersistentFlags().BoolVar(&hardR, "hard", false, "Forces VM to do a hard reset.")
	resetCmd.PersistentFlags().BoolVar(&softR, "soft", true, "Does soft stop for a vm with signal to shutdown.")
	rootCmd.AddCommand(resetCmd)
}
