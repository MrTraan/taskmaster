package tmexec

import (
	"io"
	"os"
	"time"
	"fmt"
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

func InitCmd(p []tmconf.ProcSettings) ([]ProcWrapper, error) {
	var procW	[]ProcWrapper
	var tmp		ProcWrapper

	for i, _ := range p {
		tmp.ProcSettings = p[i]
		err := tmp.initCmd()
		if err != nil {
			return nil, err
		}
		procW = append(procW, tmp)
	}
	return procW, nil
}

func (p *ProcWrapper) getStdout() error {
	if p.Stdout == "" {
		return nil
	}
	file, err := os.OpenFile(p.Stdout, os.O_CREATE | os.O_WRONLY , 0666)
	if err != nil {
		return err
	}
	p.Command.Stdout = file
	fmt.Println(p.Command.Stdout)
	return nil
}

func (p *ProcWrapper) getStderr() error {
	if p.Stderr == "" {
		return nil
	}
	file, err := os.OpenFile(p.Stderr, os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	p.Command.Stderr = file
	return nil
}

/*
** initCmd set :
** - Stdin / Stdout
** - WorkingDir
** - Environment
*/
func (p *ProcWrapper) initCmd() error {
	args := strings.Split(p.Cmd, " ")
	p.Command = exec.Command(args[0], args[1:]...)
	if err := p.getStdout(); err != nil {
		return err
	}
	if err := p.getStderr(); err != nil {
		return err
	}
	if p.WorkingDir == "" {
		p.WorkingDir = "."
	}
	p.Command.Dir = p.WorkingDir
	p.Command.Env = p.Env
	return nil
}
