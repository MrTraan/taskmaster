package tmexec

import (
	"strings"
	"vogsphere.42.fr/taskmaster.git/tmconf"
)

type ProcWrapper struct {
	tmconf.ProcSettings
	Command *exec.Cmd
	Status  string
	Time	time.Time
	Signal	syscall.Signal
}

func (p *[]ProcWrapper) InitCmd() error {
	for i, _ := range p {
		args = strings.Split(p.Cmd, " ")
		p[i].Command = exec.Command(args[0], args[1:]...)
		// need to init stdin and stdout
		// need to set the directory
		// need to set the status
		// need to set the time ??
	}
	return nil

}
