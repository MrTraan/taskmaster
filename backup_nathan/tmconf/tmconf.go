package tmconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type EnvVar struct {
	Key   string
	Value string
}

func (ev EnvVar) String() string {
	return fmt.Sprintf("%s=%s", ev.Key, ev.Value)
}

type ProcSettings struct {
	Name         string
	Cmd          string
	Numprocs     int
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

func ReadConfig(filename string) (settings []ProcSettings, err error) {
	var conf []ProcSettings

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &conf)
	return conf, err
}
