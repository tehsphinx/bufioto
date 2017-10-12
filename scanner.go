package bufioto

import (
	"bufio"
	"io"
	"log"
	"time"
)

// TimeoutScanner is a wrapper for a bufio scanner adding timeout
type TimeoutScanner struct {
	*bufio.Scanner
	timeout time.Duration
	err     error
}

// NewTimeoutScanner creates a new TimeoutScanner
func NewTimeoutScanner(reader io.Reader, timeout time.Duration) *TimeoutScanner {
	return &TimeoutScanner{
		Scanner: bufio.NewScanner(reader),
		timeout: timeout,
	}
}

// Scan wraps bufio.Scanner.Scan adding timout
func (s *TimeoutScanner) Scan() bool {
	var (
		ok   bool
		ch   = make(chan bool)
		chTo = make(chan bool, 1)
	)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Scanner panicked with:", r)
			}
		}()

		ok = s.Scanner.Scan()
		select {
		case ch <- true:
		case <-chTo:
		}
	}()

	// no timout: wait for scanner result
	if s.timeout == 0 {
		<-ch
		return ok
	}

	select {
	case <-ch:
		return ok
	case <-time.After(s.timeout):
		chTo <- true
		s.err = Timeout
		return false
	}
}

// Err wraps bufio.Scanner.Err adding timeout error
func (s *TimeoutScanner) Err() error {
	if s.err != nil {
		return s.err
	}
	return s.Scanner.Err()
}
