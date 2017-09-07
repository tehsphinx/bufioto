package bufioto

import (
	"bufio"
	"time"
)

// TimeoutScanner is a wrapper for a bufio scanner adding timeout
type TimeoutScanner struct {
	*bufio.Scanner
	timeout time.Duration
	err     error
}

// NewTimeoutScanner creates a new TimeoutScanner
func NewTimeoutScanner(scanner *bufio.Scanner, timeout time.Duration) *TimeoutScanner {
	return &TimeoutScanner{
		Scanner: scanner,
		timeout: timeout,
	}
}

// Scan wraps bufio.Scanner.Scan adding timout
func (s *TimeoutScanner) Scan() (ok bool) {
	ch := make(chan bool)
	chTo := make(chan bool, 1)
	go func() {
		ok = s.Scanner.Scan()
		select {
		case ch <- true:
		case <-chTo:
		}
	}()
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
