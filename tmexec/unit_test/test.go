package main

import (
	"vogsphere.42.fr/taskmaster.git/tmconf"
	"vogsphere.42.fr/taskmaster.git/tmexec"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "requires a config file\n")
		os.Exit(1)
	}
	conf, err := tmconf.GetProcSettings(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing: %v\n", err)
	}
	procW, err := tmexec.InitCmd(conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "InitCmd: %v\n", err)
	}
	for i, _ := range procW {
		go tmexec.StartCmd(&procW[i])
	}
	fmt.Println("new status")
	tmexec.Status(procW, "")
	fmt.Println("new status")
	tmexec.Status(procW, "lol")
	fmt.Println("new status")
	tmexec.Status(procW, "test")
	fmt.Println("new status")
	tmexec.Status(procW, " ")
}
