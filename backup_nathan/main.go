package main

import (
	"fmt"
	"github.com/MrTraan/taskmaster/tmconf"
	"gopkg.in/readline.v1"
	"os"
	"os/exec"
	"strings"
	"syscall"
    "time"
)

var completion = readline.NewPrefixCompleter(
	readline.PcItem("start"),
	readline.PcItem("stop"),
	readline.PcItem("exit"),
)

func startNewProc(conf tmconf.ProcSettings) {
    fmt.Println("My pid is ", os.Getpid())
	args := strings.Split(conf.Cmd, " ")
	path, err := exec.LookPath(args[0])
	if err != nil {
		fmt.Printf("Look path error: %v\n", err)
		return
	}
	fmt.Printf("Executing %s with args %v\n", path, args[1:])
	proc, err := os.StartProcess(path, args, &os.ProcAttr{})
	if err != nil {
		fmt.Println("Start process error: ", err)
		return
	}
	ret, err := proc.Wait()
	if err != nil {
		fmt.Println("Wait error: ", err)
		return
	}
	raw := ret.Sys()
	status, ok := raw.(syscall.WaitStatus)
	if !ok {
		fmt.Println("Conversion failure:", err)
	}
	if status.Signaled() {
		fmt.Println("Stopsignal: ", status.Signal())
	}
	fmt.Printf("Command %s exited with status %v\n", conf.Name, status.ExitStatus())
}

func cheapProc(conf tmconf.ProcSettings){
    args := strings.Split(conf.Cmd, " ")
    cmdout, err := exec.Command(args[0], args[1:]...).Output()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("command output: %s\n", cmdout)
}

func main() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "> ",
		AutoComplete: completion,
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	s, err := tmconf.ReadConfig("conf.json")
	if err != nil {
		panic(err)
	}
	//fmt.Println(s)

	for _, value := range s {
		go startNewProc(value)
	}
    time.Sleep(5 * time.Second)
	fmt.Println("done")
	
	   line := ""
	   for !strings.HasPrefix(line, "exit") {
	       line, err = rl.Readline()
	       if err != nil {
	           panic(err)
	       }
	       fmt.Println(line)
	   }
	
}
