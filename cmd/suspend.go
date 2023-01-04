package cmd

import (
	"github.com/haad/vmmanager/util"
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
			return validatePostArguments(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			var vmFlags = util.VmrunFlags{Hard: hardR, Soft: softR, Gui: false}

			util.VmrunExecCommand("suspend", args[0], &vmFlags)
		},
	}

	// suspendCmd.Flags().StringVarP(&vmx, "vmx", "V", "", "VM vmx file")
	// suspendCmd.MarkFlagRequired("vmx")
	suspendCmd.PersistentFlags().BoolVar(&hardR, "hard", false, "Forces VM to do a hard suspend.")
	suspendCmd.PersistentFlags().BoolVar(&softR, "soft", true, "Does soft stop for a vm with signal to shutdown.")
	rootCmd.AddCommand(suspendCmd)
}
