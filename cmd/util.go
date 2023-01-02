package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

type vmrunFlags struct {
	Hard bool
	Soft bool
	Gui  bool
}

func vmrunExecCommand(vmrunc string, vmx string, vmrf *vmrunFlags) {
	var vmrArgs string
	//	var vmrGArgs string

	if !vmrf.Gui {
		if vmrf.Hard {
			vmrArgs = "hard"
		} else if vmrf.Soft {
			vmrArgs = "soft"
		}
	}

	if !vmrf.Hard && !vmrf.Soft {
		if vmrf.Gui {
			vmrArgs = "gui"
		} else {
			vmrArgs = "nogui"
		}
	}

	fmt.Println(vmrf)
	fmt.Printf("vmrun %s %s %s\n", vmrunc, vmx, vmrArgs)

	// Execute the "vmrun" command in the current directory
	output, err := exec.Command("vmrun", vmrunc, vmx, vmrArgs).Output()
	if err != nil {
		fmt.Println("Error Message:", string(output))
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}

	// Print the command output
	fmt.Println(string(output))
}
