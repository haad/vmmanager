package dist

import (
	"embed"
	"io/ioutil"
	"os"

	"github.com/haad/vmmanager/log"
)

//go:embed seed.iso
var f embed.FS

func CreateSeedIso() string {
	var cdpath string

	data, _ := f.ReadFile("seed.iso")

	cdpath, err := ioutil.TempDir("", "vmmanager")
	if err != nil {
		log.Slog.Fatal(err)
	}
	// defer os.RemoveAll(cdpath)

	log.Slog.Debugf("Creating iso image at: %s/seed.iso", cdpath)
	isoF, err := os.Create(cdpath + "/seed.iso")
	if err != nil {
		log.Slog.Fatal(err)
		return ""
	}

	n, err := isoF.Write(data)
	if err != nil {
		log.Slog.Fatal(err)
		return ""
	}

	log.Slog.Debugf("Created new ISO image at %s with size: %d\n", cdpath, n)
	isoF.Sync()

	return cdpath + "/seed.iso"
}
