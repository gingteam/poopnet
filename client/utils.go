package main

import (
	"os/exec"
	"runtime"
)

func ShellExec(shellCommand string) error {
	osName := runtime.GOOS
	switch osName {
	case "windows":
		return exec.Command("cmd", "/C", shellCommand).Run()
	case "linux":
		return exec.Command("/bin/bash", "-c", shellCommand).Run()
	}
	return nil
}
