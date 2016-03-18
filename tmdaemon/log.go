package main

import (
	"os"
	"log"
)

func CreateLogFile(filename string) (*log.Logger, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	logfile := log.New(file, "taskmaster> ", log.LstdFlags)
	return logfile, nil
}
