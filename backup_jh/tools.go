package main

import (
	"os"
	"strings"
	"os/exec"
)

func initCmd(proc []ProcWrapper) {
	for i, _ := range proc {
		arg := strings.Split(proc[i].Cmd, " ")
		proc[i].Command = exec.Command(arg[0], arg[1:]...)
		proc[i].Command.Stdout = os.Stdout
		proc[i].Command.Stderr = os.Stderr
		if proc[i].WorkingDir != "" {
			proc[i].Command.Dir = proc[i].WorkingDir
		} else {
			proc[i].Command.Dir = "."
		}
		proc[i].Status = STOPPED
	}
}
