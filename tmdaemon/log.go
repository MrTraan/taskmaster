package main

import (
	"os"
	"log"
	"sync"
)

type Logfile struct {
	*log.Logger
	mutex	sync.Mutex
}

/*
** to use it in the program, just call Print on your Logile with your message
** ex: logfile.Print("This will be printed in the logfile")
** date and time will be automatically added (thamks to LstdFlags)
*/
func CreateLogFile(filename string) (*Logfile, error) {
	var logfile Logfile
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	logfile.Logger = log.New(file, "taskmaster> ", log.LstdFlags)
	return &logfile, nil
}
