package util

import (
	"os/exec"
	"path/filepath"
)

// LookPath detect the vmware command binary path.
func LookPath(cmd string) string {
	if path, err := exec.LookPath(cmd); err == nil {
		return path
	}

	return filepath.Join(cmd)
}
