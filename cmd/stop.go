package cmd

import (
	"github.com/haad/vmmanager/util"
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
			return validatePostArguments(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			var vmFlags = util.VmrunFlags{Hard: hardR, Soft: softR, Gui: false}

			util.VmrunExecCommand("stop", args[0], &vmFlags)
		},
	}

	// stopCmd.Flags().StringVarP(&vmx, "vmx", "V", "", "VM vmx file")
	// stopCmd.MarkFlagRequired("vmx")
	stopCmd.PersistentFlags().BoolVar(&hardR, "hard", false, "Forces VM to do a hard shutdown.")
	stopCmd.PersistentFlags().BoolVar(&softR, "soft", true, "Does soft stop for a vm with signal to shutdown.")
	rootCmd.AddCommand(stopCmd)
}
