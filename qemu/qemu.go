package qemu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/haad/vmmanager/log"
	"github.com/haad/vmmanager/util"
)

var qimg = util.LookPath("qemu-img")

func qemuExec(arg ...string) (string, error) {
	// vmrun with nogui on VMware Fusion through at least 8.0.1 doesn't work right
	// if the umask is set to not allow world-readable permissions
	cmd := exec.Command(qimg)
	cmd.Args = append(cmd.Args, arg...)

	log.Slog.Debugf("Executing qemu-img command: %s %s\n", cmd.Path, cmd.Args)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		if runErr := err.(*exec.ExitError); runErr != nil {
			return "", fmt.Errorf(stdout.String())
		}
	}

	return stdout.String(), err
}

func ConvertImageVmdk(spath string, dpath string) error {
	log.Slog.Debugf("Converting Source image at: %s to new vmdk image at: %s", spath, dpath)

	itype, err := getImageFormat(spath)
	if err != nil {
		log.Slog.Errorf("Couldn't get disk image format.")
		return err
	}

	// qemu-img convert -p -f qcow2 -O vmdk ubuntu-22.04-server-cloudimg-arm64.img ubuntu-22.04-server-cloudimg-arm64.vmdk
	output, err := qemuExec("convert", "-p", "-f", itype, "-O", "vmdk", spath, dpath)
	if err != nil {
		log.Slog.Errorf("Image conversion from %s to vmkdk failed, original image path: %s, destination path: %s", itype, spath, dpath)
		return fmt.Errorf("%s", output)
	}

	log.Slog.Infof(output)
	return nil
}

type QemuImgInfo struct {
	VirtualSize int64  `json:"virtual-size"`
	Filename    string `json:"filename"`
	Format      string `json:"format"`
	ActualSize  int64  `json:"actual-size"`
	DirtyFlag   bool   `json:"dirty-flag"`
}

func getImageFormat(qpath string) (string, error) {
	var qinfo QemuImgInfo
	log.Slog.Debugf("Getting image format for image: %s", qpath)

	// qemu-img info --output json alpine-aarch64-rpi-3.12.img
	output, err := qemuExec("info", "--output", "json", qpath)
	if err != nil {
		log.Slog.Errorf("Image info from image %s failed.", qpath)
		return "", fmt.Errorf("%s", output)
	}
	log.Slog.Debugf("qemu-img info output: %s", output)

	err = json.Unmarshal([]byte(output), &qinfo)
	if err != nil {
		log.Slog.Errorf("Unmarshalling json output from qemu-img failed output: %s.", output)
		return "", fmt.Errorf("Unmarshalling json output from qemu-img failed output: %s.", output)
	}

	log.Slog.Debugf("%+v", qinfo)
	return qinfo.Format, nil
}
