package deamon

import (
	"os"
	"os/exec"
	"path/filepath"
)

func Run() {
	if os.Getppid() != 1 {
		filePath, _ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()
		os.Exit(0)
	}
}
