package main

import (
	"errors"
	"fmt"
	"time"
)

type Batch []byte

type BatchSender struct {
	Delay      time.Duration
	MaxRetries int
	stats      Stats
	err        error
}

// Send sends batch to ...
func (sender *BatchSender) Send(b Batch) bool {
	if sender.err != nil {
		return true
	}
	// test logic that may return error.
	startTime := time.Now()
	sender.stats.RetryN++
	err := testCall(sender.stats.RetryN)
	if err == nil {
		sender.stats.ExecTime += time.Since(startTime)
		return false
	}
	if sender.stats.RetryN < sender.MaxRetries {
		<-time.Tick(sender.Delay)
		sender.stats.ExecTime += time.Since(startTime)
		return true
	}
	sender.err = err
	sender.stats.ExecTime += time.Since(startTime)
	return false
}

// Err returns occured error while sending.
func (sender *BatchSender) Err() error {
	return sender.err
}

// Stats returns collected metrics.
func (sender *BatchSender) Stats() *Stats {
	return &sender.stats
}

type Stats struct {
	RetryN   int
	ExecTime time.Duration
}

func (stats *Stats) Reset() {
	stats.RetryN = 0
	stats.ExecTime = 0
}

// testCall returns success (nil) only after 4-th attempt.
func testCall(n int) error {
	logger.Debug(fmt.Sprintf("iteration: %d", n))
	if n > 4 {
		logger.Debug("success")
		return nil
	}
	return errors.New("error occured")
}
