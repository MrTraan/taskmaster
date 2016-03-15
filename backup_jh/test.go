package main

import (
	"os"
	"fmt"
	"os/exec"
	"io/ioutil"
	"strings"
	"encoding/json"
)

type Process struct {
	Name			string `json:Name`
	Cmd				string `json:cmd`
	Numprocs		int `json:numprocs`
	Umask			int `json:umask`
	Workingdir		string `json:workingdir`
	Autostart		bool `json:autostart`
	Autorestart		string `json:autorestart`
	Exitcodes		[]int `json:exitcodes`
	Startretries	int `json:startretries`
	Starttime		int `json:starttime`
	Stoptime		int `json:stoptime`
	Stdout			string `json:stdout`
	Stderr			string `json:stderr`
}

type Container struct {
	Programs []Process
}

func main() {
	var container Container
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "taskmaster: requires a config file\n")
		return
	}
	config_file := os.Args[1]
	data := parseFile(config_file)
	err := json.Unmarshal(data, &container)
	if err != nil {
		fmt.Fprintf(os.Stderr, "critical error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
	for i := 0; i < len(container.Programs); i++ {
		fmt.Println(container.Programs[i])
	}
	//Test execution
	for _, v := range container.Programs {
		tmp := strings.Split(v.Cmd, " ")
		fmt.Println(tmp)
		cmd := exec.Command(tmp[0], tmp[1:]...)
		result, err := cmd.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		} else {
			fmt.Println("Cmd successfully launched:")
			fmt.Println(string(result))
		}
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

func (p Process) String() string {
	str := fmt.Sprintf("Name: %s -> cmd %s\n", p.Name, p.Cmd)
	str += fmt.Sprintf("\tNProcs: %d || Umask: %d\n", p.Numprocs, p.Umask)
	str += fmt.Sprintf("\tOut: %s || Err: %s\n", p.Stdout, p.Stderr)
	str += fmt.Sprintf("\tWDir: %s || Autorestart: %t\n", p.Workingdir, p.Autostart)
	return str
}
