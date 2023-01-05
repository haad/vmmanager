package cmd

import (
	"github.com/haad/vmmanager/vmware"
	"github.com/spf13/cobra"
)

func init() {
	var wait bool
	var getIpCmd = &cobra.Command{
		Use:   "getip [path_to_file.vmx]",
		Short: "getip virtual machine ",
		Long:  "Try to guess host IP from vmware tools.",
		Args: func(cmd *cobra.Command, args []string) error {
			return validatePostArguments(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			f := vmware.NewFusion(args[0], "", "", "")
			f.GetGuestIPAddress(wait)
		},
	}

	getIpCmd.PersistentFlags().BoolVar(&wait, "wait", false, "Wait for results.")
	rootCmd.AddCommand(getIpCmd)
}
