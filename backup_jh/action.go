package main

import (
	"fmt"
	"vogsphere.42.fr/taskmaster.git/backup_nathan/tmconf"
)

func status(proc []ProcWrapper) {
	for _, v := range proc {
		fmt.Printf("%s: %s -> status %s\n", v.Name, v.Cmd, v.Status)
	}
}

func reload(filename string, proc []ProcWrapper) ([]tmconf.ProcSettings, error) {
	conf, err := tmconf.ReadConfig(filename)
	if err != nil {
		return nil, err
	}
	// for i, _ := range proc {
	// 	// if proc[i].ProcSettings != conf[i] {
	// 	// 	//fmt.Printf("%s changed during reload", proc[i].Cmd)
	// 	// }
	// }
	return conf, nil
}

func start(proc tmconf.ProcSettings) {
	fmt.Println("gonna start a new proc")
	fmt.Println(proc)
}

func stop(proc tmconf.ProcSettings) {
	fmt.Println("gonna stop a proc")
	fmt.Println(proc)
}
