package cmd

import (
	"github.com/haad/vmmanager/vmware"
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
			f := vmware.NewFusion(args[0], "", "", "")
			f.Start(gui)
		},
	}

	startCmd.PersistentFlags().BoolVar(&gui, "gui", false, "Enables gui window for vm start.")
	//startCmd.Flags().StringVarP(&vmx, "vmx", "V", "", "VM vmx file")
	//startCmd.MarkFlagRequired("vmx")
	rootCmd.AddCommand(startCmd)

}
