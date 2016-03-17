package tmconf

import (
	"encoding/json"
	"fmt"
    "io/ioutil"
)

type ProcSettings struct {
	Name         string
	Cmd          string
	Numprocs     int
	WorkingDir	 string
	Autostart    bool
	Autorestart  string
	Exitcodes    []int
	Startretries int
	Starttime    int
	Stopsignal   string
	Stoptime     int
	Stdout       string
	Stderr       string
	Env          []string
}

func (s ProcSettings) String() string {
	return fmt.Sprintf("{Name: %s\nCmd: %s\nNumprocs: %d\nAutostart: %v\nAutorestart: %s\nExitcode: %v\nStartretries: %d\nStarttime: %d\nStopsignal: %s\nStoptime: %d\nStdout: %s\nStderr: %s\nEnv: %v}\n",
		s.Name, s.Cmd, s.Numprocs, s.Autostart, s.Autorestart, s.Exitcodes, s.Startretries,
		s.Starttime, s.Stopsignal, s.Stoptime, s.Stdout, s.Stderr, s.Env)
}

func GetProcSettings(filename string) (settings []ProcSettings, err error) {
	conf, err := ReadConfig(filename)
	if err != nil {
		return nil, err
	}
	for i, _ := range conf {
		if conf[i].Numprocs != 1 {
			for j := 1; j < conf[i].Numprocs; j++ {
				conf = append(conf, conf[i])
			}
		}
	}
	return conf, nil
}

func ReadConfig(filename string) (settings []ProcSettings, err error) {
	var conf []ProcSettings

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &conf)
	return conf, err
}
