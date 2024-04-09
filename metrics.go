package main

import (
	"log"
	"time"
)

type Metrics struct {
}

func (metr *Metrics) ExecutionTime(caller string, d time.Duration) {
	log.Printf("METRICS: caller=%s, execution_time=%s\n", caller, d)
}

func (metr *Metrics) RetryCount(caller string, n int) {
	log.Printf("METRICS: caller=%s, retry_count=%d\n", caller, n)
}

func (metr *Metrics) ErrorsIncr(caller string) {
	log.Printf("METRICS: caller=%s, errors_count=<n>\n", caller)
}
