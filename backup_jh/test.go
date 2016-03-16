package main

import (
	"fmt"
	"gopkg.in/readline.v1"
	"os"
	"log"
	"os/exec"
	"strings"
	"time"
	"syscall"
	"vogsphere.42.fr/taskmaster.git/backup_nathan/tmconf"
)

const (
	STARTING = "Starting"
	RUNNING  = "Running"
	STOPPED  = "Stopped"
	FINISHED = "Finished"
)

type ProcWrapper struct {
	tmconf.ProcSettings
	Command *exec.Cmd
	Status  string
	Time	time.Time
	Signal	syscall.Signal
}

var completion = readline.NewPrefixCompleter(
	readline.PcItem("start"),
	readline.PcItem("stop"),
	readline.PcItem("exit"),
	readline.PcItem("reload"),
	readline.PcItem("status"),
)

func main() {
	var container []ProcWrapper
	var tmp_procw ProcWrapper

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "taskmaster: requires a config file or an instruction\n")
		return
	}
	log, err := testLogFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "taskmaster: log error: %v\n", err)
		os.Exit(1)
	}
	initCmd(container)
	testExec(container, log)
	config_file := os.Args[1]
	tmp, err := tmconf.ReadConfig(config_file)
	for _, v := range tmp {
		tmp_procw.ProcSettings = v
		tmp_procw.Status = STOPPED
		container = append(container, tmp_procw)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "taskmaster: %v\n", err)
		os.Exit(1)
	}
	testPrompt(container)
}


func testPrompt(proc []ProcWrapper) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "> ",
		AutoComplete: completion,
	})
	if err != nil {
		fmt.Println("couldn't launch readline, program aborted")
		os.Exit(1)
	}
	for line, err := "", error(nil); !strings.HasPrefix(line, "exit"); {
		line, err = rl.Readline()
		if err != nil {
			panic(err)
		}
		if strings.HasPrefix(line, "status") {
			status(proc)
		}
	}
}

func testLogFile() (*log.Logger, error) {
	file, err := os.Create("log.txt")
	if err != nil {
		return nil, err
	}
	log := log.New(file, "taskmaster> ", log.LstdFlags)
	return log, nil
}

func lolTest() int {
	return 4
}

func testArrayFunc() []func() int {
	test := []func() int{}
	test = append(test, func() int { return 1 })
	test = append(test, func() int { return 2 })
	test = append(test, func() int { return 3 })
	test = append(test, lolTest)
	for _, v := range test {
		fmt.Println(v())
	}
	return test
}

func testExec(proc []ProcWrapper, log *log.Logger) {
	for _, v := range proc {
		log.Output(2, fmt.Sprintf("launching %s in %s\n", v.Cmd, v.Command.Dir))
		err := v.Command.Start()
		if err != nil {
			fmt.Println(v)
			continue
		}
		v.Status = RUNNING
		fmt.Printf("Status of %s: %s\n", v.Cmd, v.Status)
		fmt.Println("Waiting")
		err = v.Command.Wait()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s failed: %v\n", v.Cmd, err)
		}
		v.Status = STOPPED
		fmt.Printf("Status of %s: %s\n", v.Cmd, v.Status)
		exit_value := v.Command.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
		fmt.Println("\033[31mExit value:\033[0m", exit_value)
		fmt.Println("Done")
	}
}
