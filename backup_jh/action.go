package main

import (
	"fmt"
	"vogsphere.42.fr/taskmaster.git/backup_nathan/tmconf"
)

func status(proc tmconf.ProcSettings) {
	fmt.Println("gonna print the status")
	fmt.Println(proc)
}

func reload(filename string) {
	fmt.Println("gonna reload a new file")
	fmt.Println(filename)
}

func start(proc tmconf.ProcSettings) {
	fmt.Println("gonna start a new proc")
	fmt.Println(proc)
}

func stop(proc tmconf.ProcSettings) {
	fmt.Println("gonna stop a proc")
	fmt.Println(proc)
}
