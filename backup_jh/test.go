package main

import (
	"os"
	"fmt"
	"strings"
	"os/exec"
	"syscall"
	"gopkg.in/readline.v1"
	"vogsphere.42.fr/taskmaster.git/backup_nathan/tmconf"
)

var completion = readline.NewPrefixCompleter(
	readline.PcItem("start"),
	readline.PcItem("stop"),
	readline.PcItem("exit"),
	readline.PcItem("reload"),
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "taskmaster: requires a config file\n")
		return
	}
	rl, err := readline.NewEx(&readline.Config{
		Prompt: "> ",
		AutoComplete: completion,
	})
	for line, err := "", error(nil); !strings.HasPrefix(line, "exit"); {
		line, err = rl.Readline()
		if err != nil {
			panic(err)
		}
		fmt.Println(line)
	}
	config_file := os.Args[1]
	container, err := tmconf.ReadConfig(config_file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "taskmaster: %v\n", err)
		os.Exit(1)
	}
	testExec(container)
}

func testExec(proc []tmconf.ProcSettings) {
	for i, v := range proc {
		cmd_splitted := strings.Split(v.Cmd, " ")
		cmd := exec.Command(cmd_splitted[0], cmd_splitted[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = proc[i].WorkingDir
		fmt.Printf("launching %s in %s\n", proc[i].Cmd, cmd.Dir)
		err := cmd.Start()
		if err != nil {
			fmt.Println(v)
			continue
		}
		fmt.Println("Waiting")
		err = cmd.Wait()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s failed: %v\n", v.Cmd, err)
		}
		fmt.Println(cmd.ProcessState.String())
		exit_value := cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
		fmt.Println("\033[31mExit value:\033[0m", exit_value)
		fmt.Println("Done")
	}
}
