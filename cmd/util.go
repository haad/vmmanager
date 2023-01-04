package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func validatePostArguments(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
	}

	if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
		return err
	}

	parts := strings.Split(args[0], ".")
	if strings.Contains(parts[len(parts)-1], "vmx") {
		return nil
	}

	return fmt.Errorf("invalid vmx path specified: %s", args[0])
}
