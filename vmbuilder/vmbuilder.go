package vmbuilder

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/haad/vmmanager/dist"
	"github.com/haad/vmmanager/log"
	"github.com/haad/vmmanager/qemu"
)

type VM struct {
	name string
	vmx  string

	baseDir string
	vmDir   string

	// Required VM resources
	cpu       uint16
	mem       uint16
	disk      uint16
	diskPath  string
	cdromPath string

	sourceImage string
}

func NewVM(path, vmx, name, dir string, cpu, mem, disk uint16) *VM {
	return &VM{
		name: name,
		vmx:  fmt.Sprintf("%s/%s.vmwarevm/%s.vmx", dir, name, name), // full Path to a generated VMX file.

		baseDir: dir,
		vmDir:   fmt.Sprintf("%s/%s.vmwarevm", dir, name), // Path to VM directory with vmdk + vmx files.

		cpu:       cpu,
		mem:       mem,
		disk:      disk,
		diskPath:  path,
		cdromPath: dist.CreateSeedIso(),

		sourceImage: path,
	}
}

func (v *VM) BuildVm() error {
	log.Slog.Infof("Building vm at %s, name: %s, vmx: %s, diskPath: %s", v.baseDir, v.name, v.vmx, v.diskPath)

	checkVMDir(v.vmDir)
	qemu.ConvertImageVmdk(v.sourceImage, v.getVMDKpath())
	_, err := v.VmxRender()
	if err != nil {
		log.Slog.Errorf("Falied to create VMX template.")
		return err
	}

	return nil
}

func (v *VM) CleanUp() {
	os.RemoveAll(v.cdromPath)
}

func (v *VM) getVMDKpath() string {
	file := filepath.Base(v.diskPath)

	// XXX: Hack ?
	v.diskPath = fmt.Sprintf("%s/%s.vmdk", v.vmDir, file[:len(file)-len(filepath.Ext(file))])

	return v.diskPath
}

func checkVMDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Slog.Debugf("Creating directory for vm at: %s", dir)
		os.MkdirAll(dir, 0750)
	}
}
