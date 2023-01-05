package cmd

import (
	"os"

	"github.com/haad/vmmanager/log"
	"github.com/haad/vmmanager/util"
	"github.com/haad/vmmanager/vmbuilder"
	"github.com/spf13/cobra"
)

func init() {
	var url, fpath, vmx, name, path, dir string
	var cpu, mem, disk uint16
	var err error

	var buildCmd = &cobra.Command{
		Use:   "build",
		Short: "Build a VM from a QCOW2 image",
		Long:  "Build a VM from a QCOW2 image",
		Run: func(cmd *cobra.Command, args []string) {
			if url != "" {
				log.Slog.Infof("Downloading image from: %s", url)

				path, err = util.DownloadHttpFile(url)
				if err != nil {
					log.Slog.Errorf("Error saving file: %s", err)
				}
				log.Slog.Infof("Disk image downloaded to: %s", path)
			} else if fpath != "" {
				log.Slog.Infof("Disk image present at: %s", fpath)
				path = fpath
			} else {
				log.Slog.Errorf("url or path flag required.")
				os.Exit(1)
			}

			vm := vmbuilder.NewVM(path, vmx, name, dir, cpu, mem, disk)
			vm.BuildVm()
		},
	}

	buildCmd.Flags().StringVarP(&name, "name", "n", "", "vm name")
	buildCmd.Flags().StringVarP(&url, "url", "u", "", "URL to fetch qcow2 image from.")
	buildCmd.Flags().StringVarP(&dir, "dir", "b", "", "Direcotry where to place your vms.")
	buildCmd.Flags().StringVarP(&fpath, "path", "f", "", "Path to a local file used as source image.")
	buildCmd.Flags().StringVarP(&vmx, "vmx", "v", "", "vmx file dest path.")

	buildCmd.Flags().Uint16VarP(&cpu, "cpu", "c", 1, "Number of cores allocated to vm")
	buildCmd.Flags().Uint16VarP(&mem, "mem", "m", 2048, "Memory allocated for vm in megabytes")
	buildCmd.Flags().Uint16VarP(&disk, "disk", "d", 10, "Disk allocated for vm in gigabytes")

	buildCmd.MarkFlagRequired("name")
	//buildCmd.MarkFlagRequired("vmx")

	rootCmd.AddCommand(buildCmd)
}
