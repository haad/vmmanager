package cmd

import (
	"fmt"

	log "github.com/haad/vmmanager/log"
	"github.com/haad/vmmanager/util"
	"github.com/spf13/cobra"
)

func init() {
	var url, vmx, name string
	var cpu, mem, disk, int16
	var buildCmd = &cobra.Command{
		Use:   "build",
		Short: "Build a VM from a QCOW2 image",
		Long:  "Build a VM from a QCOW2 image",
		Run: func(cmd *cobra.Command, args []string) {
			log.Slog.Infof("Downloading image from: %s", url)

			path, err := util.DownloadHttpFile(url)
			if err != nil {
				log.Slog.Errorf("Error saving file: %s", err)
			}

			log.Slog.Infof("Disk image downloaded to: %s", path)
			util.ConvertQcow2Vmdk(path, fmt.Sprintf("%s.vmdk", path))
		},
	}

	buildCmd.Flags().StringVarP(&name, "name", "n", "", "vm name")
	buildCmd.Flags().StringVarP(&url, "url", "u", "", "URL to fetch qcow2 image from.")
	buildCmd.Flags().StringVarP(&vmx, "vmx", "v", "", "vmx file dest path.")
	buildCmd.Flags().Int16VarP(&cpu, "cpu", "c", 1, "Number of cores allocated to vm")
	buildCmd.Flags().Int16VarP(&mem, "mem", "m", 2048, "Memory allocated for vm in megabytes")
	buildCmd.Flags().Int16VarP(&disk, "disk", "d", 10, "Disk allocated for vm in gigabytes")

	buildCmd.MarkFlagRequired("url")
	buildCmd.MarkFlagRequired("name")
	buildCmd.MarkFlagRequired("vmx")

	rootCmd.AddCommand(buildCmd)
}
