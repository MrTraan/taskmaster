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
	c := 0
	for _, v := range procW {
		if v.Command == nil || v.StdoutPipe == nil || v.StderrPipe == nil {
			c += 1
		}
	}
	if c != 0 {
		fmt.Printf("\033[31m%d cmds are unset\033[0m\n")
	} else {
		fmt.Println("\033[32mall process cmds are well set\033[0m")
	}
}