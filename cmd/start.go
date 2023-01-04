package cmd

import (
	"github.com/haad/vmmanager/util"
	"github.com/spf13/cobra"
)

func init() {
	var gui bool
	var startCmd = &cobra.Command{
		Use:   "start [--gui] [path_to_file.vmx]",
		Short: "Start virtual machine ",
		Long:  "This command starts vmware fusion virtual machine.",
		Args: func(cmd *cobra.Command, args []string) error {
			return validatePostArguments(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			var vmFlags = util.VmrunFlags{Hard: false, Soft: false, Gui: gui}

			util.VmrunExecCommand("start", args[0], &vmFlags)
		},
	}

	startCmd.PersistentFlags().BoolVar(&gui, "gui", false, "Enables gui window for vm start.")
	//startCmd.Flags().StringVarP(&vmx, "vmx", "V", "", "VM vmx file")
	//startCmd.MarkFlagRequired("vmx")
	rootCmd.AddCommand(startCmd)

}
