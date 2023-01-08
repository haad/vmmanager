package vmbuilder

import (
	"os"

	"github.com/flosch/pongo2"
	"github.com/haad/vmmanager/log"
)

var vmxTemplate = `
.encoding = "UTF-8"
config.version = "8"
virtualHW.version = "20"
guestOS = "arm-ubuntu-64"
guestInfo.detailed.data = "architecture='Arm' bitness='64' distroName='Ubuntu 22.04.1 LTS' distroVersion='22.04' familyName='Linux' kernelVersion='5.15.0-56-generic' prettyName='Ubuntu 22.04.1 LTS'"
tools.syncTime = "TRUE"
tools.upgrade.policy = "upgradeAtPowerCycle"
cleanShutdown = "FALSE"
softPowerOff = "FALSE"
firmware = "efi"
bios.bootOrder = "HDD"
bios.hddOrder = "nvme0:0"
displayName = "{{ name }}"
nvram = "{{ name }}.nvram"
extendedConfigFile = "{{ name }}.vmxf"
vmxstats.filename = "{{ name }}.scoreboard"
virtualHW.productCompatibility = "hosted"
cpuid.coresPerSocket = "{{ cpus }}"
memsize = "{{ memory }}"
numa.autosize.cookie = "10012"
numa.autosize.vcpu.maxPerVirtualNode = "1"
vmci0.present = "TRUE"
vmci0.id = "-1859122366"
hpet0.present = "TRUE"
svga.vramSize = "268435456"
svga.guestBackedPrimaryAware = "TRUE"
pciBridge0.pciSlotNumber = "17"
pciBridge4.pciSlotNumber = "21"
pciBridge5.pciSlotNumber = "22"
pciBridge6.pciSlotNumber = "23"
pciBridge7.pciSlotNumber = "24"
pciBridge0.present = "TRUE"
pciBridge4.present = "TRUE"
pciBridge4.virtualDev = "pcieRootPort"
pciBridge4.functions = "8"
pciBridge5.present = "TRUE"
pciBridge5.virtualDev = "pcieRootPort"
pciBridge5.functions = "8"
pciBridge6.present = "TRUE"
pciBridge6.virtualDev = "pcieRootPort"
pciBridge6.functions = "8"
pciBridge7.present = "TRUE"
pciBridge7.virtualDev = "pcieRootPort"
pciBridge7.functions = "8"
sata0.pciSlotNumber = "34"
sata0.present = "TRUE"
sata0:1.deviceType = "cdrom-image"
sata0:1.fileName = "{{ seedISOimage }}"
sata0:1.present = "TRUE"
nvme0.pciSlotNumber = "224"
nvme0:0.redo = ""
nvme0.subnqnUUID = "52 4e 36 93 ac 00 66 b6-9b 7d 2f 2d 7c c6 fe e2"
nvme0.present = "TRUE"
nvme0:0.fileName = "{{ diskPath }}"
nvme0:0.present = "TRUE"
ethernet0.connectionType = "{{ netType }}"
ethernet0.addressType = "generated"
ethernet0.virtualDev = "e1000e"
ethernet0.linkStatePropagation.enable = "TRUE"
ethernet0.present = "TRUE"
ethernet0.generatedAddress = "00:0c:29:2c:c3:03"
ethernet0.generatedAddressOffset = "0"
ethernet0.pciSlotNumber = "160"
ethernet0.startConnected = "FALSE"
monitor.phys_bits_used = "36"
usb.present = "TRUE"
ehci.present = "TRUE"
usb_xhci.present = "TRUE"
usb.pciSlotNumber = "32"
ehci.pciSlotNumber = "33"
usb_xhci.pciSlotNumber = "192"
usb_xhci:4.present = "TRUE"
usb_xhci:4.deviceType = "hid"
usb_xhci:4.port = "4"
usb_xhci:4.parent = "-1"
usb_xhci:6.speed = "2"
usb_xhci:6.present = "TRUE"
usb_xhci:6.deviceType = "hub"
usb_xhci:6.port = "6"
usb_xhci:6.parent = "-1"
usb_xhci:7.speed = "4"
usb_xhci:7.present = "TRUE"
usb_xhci:7.deviceType = "hub"
usb_xhci:7.port = "7"
usb_xhci:7.parent = "-1"
usb_xhci:4.present = "TRUE"
usb_xhci:4.deviceType = "hid"
usb_xhci:4.port = "4"
usb_xhci:4.parent = "-1"

floppy0.present = "FALSE"
`

// type VMx struct {
// 	Encoding     string `vmx:".encoding"`
// 	Annotation   string `vmx:"annotation"`
// 	Hwversion    uint8  `vmx:"virtualHW.version"`
// 	HwProdCompat string `vmx:"virtualHW.productCompatibility"`
// 	Memsize      uint   `vmx:"memsize"`
// 	Numvcpus     uint   `vmx:"numvcpus"`
// 	MemHotAdd    bool   `vmx:"mem.hotadd"`
// 	DisplayName  string `vmx:"displayName"`
// 	GuestOS      string `vmx:"guestOS"`
// 	Autoanswer   bool   `vmx:"msg.autoAnswer"`
// }

func (v *VM) VmxRender() (string, error) {
	var f *os.File
	// Compile the template first (i. e. creating the AST)
	tpl, err := pongo2.FromString(vmxTemplate)
	if err != nil {
		log.Slog.Errorf("Compiling VMX template for VM: %s failed.", v.name)
		return "", err
	}
	// Now you can render the template with the given
	// pongo2.Context how often you want to.
	out, err := tpl.Execute(pongo2.Context{"name": v.name, "memory": v.mem, "cpus": v.cpu, "diskPath": v.diskPath, "seedISOimage": v.cdromPath, "netType": "nat"})
	if err != nil {
		log.Slog.Errorf("Rendering VMX template for VM: %s failed.", v.name)
		return "", err
	}
	log.Slog.Debugf(out)

	// Create VMX file on destination path v.vmx
	f, err = os.Create(v.vmx)
	if err != nil {
		log.Slog.Error(err)
		return "", err
	}

	log.Slog.Debugf("Creating VMX file at: %s", v.vmx)
	_, err = f.WriteString(out + "\n")
	if err != nil {
		log.Slog.Error(err)
		return "", err
	}

	f.Sync()

	return out, err
}
