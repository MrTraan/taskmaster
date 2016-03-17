package tmexec

import (
	"io"
	"time"
	"syscall"
	"strings"
	"os/exec"
	"vogsphere.42.fr/taskmaster.git/tmconf"
)

type ProcWrapper struct {
	tmconf.ProcSettings
	Command		*exec.Cmd
	StdoutPipe	io.ReadCloser
	StderrPipe	io.ReadCloser
	Status  	string
	Time		time.Time
	Signal		syscall.Signal
}

func InitCmd(p []ProcWrapper) error {
	for i, _ := range p {
		err := p[i].initCmd()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p ProcWrapper) getStdout() error {
	var err error

	if p.Stdout != "" {
		p.StdoutPipe, err = p.Command.StdoutPipe()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p ProcWrapper) getStderr() error {
	var err error

	if p.Stderr != "" {
		p.StderrPipe, err = p.Command.StderrPipe()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p ProcWrapper) initCmd() error {
	args := strings.Split(p.Cmd, " ")
	p.Command = exec.Command(args[0], args[1:]...)
	if err := p.getStdout(); err != nil {
		return err
	}
	if err := p.getStderr(); err != nil {
		return err
	}
	// need to init stdin and stdout
	// need to set the directory
	// need to set the status
	// need to set the time ??
	return nil
}
