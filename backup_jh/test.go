package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"vogshpere.42.fr/taskmaster.git/backup_nathan/tmconf"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "taskmaster: requires a config file\n")
		return
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
	fmt.Println(len(proc), "procs found\n")
	for _, v := range  proc {
		fmt.Println(v)
	}
}

func parseFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "taskmaster: %v\n", err)
		os.Exit(1)
	}
	return data
}
