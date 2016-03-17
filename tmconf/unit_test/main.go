package main

import (
	"fmt"
	"os"
	"vogsphere.42.fr/taskmaster.git/tmconf"
)

func countProc(proc []tmconf.ProcSettings) {
	var mapProc = make(map[string]int)

	for _, v := range proc {
		mapProc[v.Name] += 1
	}
	for k, v := range mapProc {
		fmt.Printf("%s : %d\n", k, v)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "requires a config file\n")
		os.Exit(1)
	}
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("\n\033[31mParsing\033[0m", os.Args[i])
		conf, err := tmconf.GetProcSettings(os.Args[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		countProc(conf)
	}

}
