package main

import "log"

type Logger struct {
}

func (l *Logger) Error(err error) {
	log.Println("ERROR:", err)
}

func (l *Logger) Debug(msg string) {
	log.Println("DEBUG:", msg)
}
