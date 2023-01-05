package cmd

import (
	"github.com/haad/vmmanager/vmware"
	"github.com/spf13/cobra"
)

func init() {
	var hardR bool
	var stopCmd = &cobra.Command{
		Use:   "stop [--hard] [path_to_file.vmx]",
		Short: "Stop virtual machine ",
		Long:  "This command stops vmware fusion virtual machine.",
		Args: func(cmd *cobra.Command, args []string) error {
			return validatePostArguments(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			f := vmware.NewFusion(args[0], "", "", "")
			if hardR {
				f.ShutDown()
			} else {
				f.Halt()
			}
		},
	}

	// stopCmd.Flags().StringVarP(&vmx, "vmx", "V", "", "VM vmx file")
	// stopCmd.MarkFlagRequired("vmx")
	stopCmd.PersistentFlags().BoolVar(&hardR, "hard", false, "Forces VM to do a hard shutdown.")
	rootCmd.AddCommand(stopCmd)
}
