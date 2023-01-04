package util

import (
	"os/exec"

	"github.com/haad/vmmanager/log"
)

var qimg = "qemu-img"

func ConvertQcow2Vmdk(qpath string, vpath string) error {
	log.Slog.Debugf("Converting Qcow2 image at: %s to new vmdk image at: %s", qpath, vpath)

	path, err := exec.LookPath(qimg)
	if err != nil {
		log.Slog.Errorf("didn't find %s executable\n", qimg)
		return err
	}

	output, err := exec.Command(path).Output()
	if err != nil {
		log.Slog.Errorf("Error Message:", string(output))
		log.Slog.Errorf("Error executing command:", err)
		return err
	}

	log.Slog.Infof(string(output))
	return nil
}
