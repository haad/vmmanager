package util

import (
	"os"
	"os/exec"

	log "github.com/haad/vmmanager/log"
)

type VmrunFlags struct {
	Hard bool
	Soft bool
	Gui  bool
}

func VmrunExecCommand(vmrunc string, vmx string, vmrf *VmrunFlags) {
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

	//fmt.Println(vmrf)
	log.Slog.Debugf("Executing vmrun command: vmrun %s %s %s\n", vmrunc, vmx, vmrArgs)

	// Execute the "vmrun" command in the current directory
	output, err := exec.Command("vmrun", vmrunc, vmx, vmrArgs).Output()
	if err != nil {
		log.Slog.Errorf("Error Message:", string(output))
		log.Slog.Errorf("Error executing command:", err)
		os.Exit(1)
	}

	log.Slog.Infof(string(output))
}
