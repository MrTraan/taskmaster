package main

import (
	"os"
	"fmt"
	"vogsphere.42.fr/taskmaster.git/taskmaster/tmconf"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "requires a config file\n")
		os.Exit(1)
	}
	conf, err := tmconf.ReadConfig(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	for _, v := range conf {
		fmt.Println(v)
	}
}
