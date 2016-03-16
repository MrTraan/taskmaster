package main

import (
	"fmt"
	"gopkg.in/readline.v1"
	"os"
	"os/exec"
	"strings"
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
}

var completion = readline.NewPrefixCompleter(
	readline.PcItem("start"),
	readline.PcItem("stop"),
	readline.PcItem("exit"),
	readline.PcItem("reload"),
	readline.PcItem("status"),
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "taskmaster: requires a config file\n")
		return
	}
	if strings.Compare(os.Args[1], "prompt") == 0 {
		fmt.Println("test the prompt")
		testPrompt()
	} else {
		config_file := os.Args[1]
		tmp, err := tmconf.ReadConfig(config_file)

		var container []ProcWrapper
		var test ProcWrapper
		for _, v := range tmp {
			test.ProcSettings = v
			test.Status = STOPPED
			container = append(container, test)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "taskmaster: %v\n", err)
			os.Exit(1)
		}
		initCmd(container)
		testExec(container)
	}
}

func testPrompt() {
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
	}
}

func testArrayFunc() []func() int {
	test := []func() int{}
	test = append(test, func() int { return 1 })
	test = append(test, func() int { return 2 })
	test = append(test, func() int { return 3 })
	fmt.Println(test[0]())
	fmt.Println(test[1]())
	fmt.Println(test[2]())
	return test
}

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

func testExec(proc []ProcWrapper) {
	for _, v := range proc {
		fmt.Printf("launching %s in %s\n", v.Cmd, v.Command.Dir)
		err := v.Command.Start()
		if err != nil {
			fmt.Println(v)
			continue
		}
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
