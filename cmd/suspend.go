package cmd

import (
	"github.com/haad/vmmanager/vmware"
	"github.com/spf13/cobra"
)

func init() {
	var hardR bool
	var suspendCmd = &cobra.Command{
		Use:   "suspend [--hard] [path_to_file.vmx]",
		Short: "Suspends virtual machine ",
		Long:  "This command suspends vmware fusion virtual machine.",
		Args: func(cmd *cobra.Command, args []string) error {
			return validatePostArguments(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			f := vmware.NewFusion(args[0], "", "", "")

			f.Suspend(hardR)
		},
	}

	// suspendCmd.Flags().StringVarP(&vmx, "vmx", "V", "", "VM vmx file")
	// suspendCmd.MarkFlagRequired("vmx")
	suspendCmd.PersistentFlags().BoolVar(&hardR, "hard", false, "Forces VM to do a hard suspend.")
	rootCmd.AddCommand(suspendCmd)
}
