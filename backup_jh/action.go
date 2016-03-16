package main

import (
	"fmt"
	"vogsphere.42.fr/taskmaster.git/backup_nathan/tmconf"
)

func status(proc []ProcWrapper) {
	for _, v := range proc {
		fmt.Println("%s -> status %s\n", v.Cmd, v.Status)
	}
}

func reload(filename string) {
}

func start(proc tmconf.ProcSettings) {
	fmt.Println("gonna start a new proc")
	fmt.Println(proc)
}

func stop(proc tmconf.ProcSettings) {
	fmt.Println("gonna stop a proc")
	fmt.Println(proc)
}
