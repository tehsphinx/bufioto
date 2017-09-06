package bufioto

import (
	"bufio"
	"errors"
	"time"
)

var Timeout = errors.New("Timeout on reading")

// TimeoutReader is a wrapper for any reader adding timeout
type TimeoutReader struct {
	bufio.Reader
	timeout time.Duration
	ch      chan bool
}

// NewTimeoutReader creates a new TimeoutReader
func NewTimeoutReader(reader bufio.Reader, timeout time.Duration) *TimeoutReader {
	return &TimeoutReader{
		Reader:  reader,
		timeout: timeout,
		ch:      make(chan bool),
	}
}

// Peek wraps bufio.Peek adding timeout
func (s *TimeoutReader) Peek(n int) (b []byte, err error) {
	go func() {
		b, err = s.Reader.Peek(n)
		s.ch <- true
	}()
	select {
	case <-s.ch:
		return
	case <-time.After(s.timeout):
		return nil, Timeout
	}
	return
}

// Discard wraps bufio.Discard adding timeout
func (s *TimeoutReader) Discard(n int) (discarded int, err error) {
	go func() {
		discarded, err = s.Reader.Discard(n)
		s.ch <- true
	}()
	select {
	case <-s.ch:
		return
	case <-time.After(s.timeout):
		return 0, Timeout
	}
	return
}

// Read wraps bufio.Read adding timeout
func (s *TimeoutReader) Read(buf []byte) (n int, err error) {
	go func() {
		n, err = s.Reader.Read(buf)
		s.ch <- true
	}()
	select {
	case <-s.ch:
		return
	case <-time.After(s.timeout):
		return 0, Timeout
	}
	return
}

// ReadByte wraps bufio.ReadByte adding timeout
func (s *TimeoutReader) ReadByte() (b byte, err error) {
	go func() {
		b, err = s.Reader.ReadByte()
		s.ch <- true
	}()
	select {
	case <-s.ch:
		return
	case <-time.After(s.timeout):
		return 0, Timeout
	}
	return
}

// ReadRune wraps bufio.ReadRune adding timeout
func (s *TimeoutReader) ReadRune() (r rune, size int, err error) {
	go func() {
		r, size, err = s.Reader.ReadRune()
		s.ch <- true
	}()
	select {
	case <-s.ch:
		return
	case <-time.After(s.timeout):
		return 0, 0, Timeout
	}
	return
}

// ReadSlice wraps bufio.ReadSlice adding timeout
func (s *TimeoutReader) ReadSlice(delim byte) (line []byte, err error) {
	go func() {
		line, err = s.Reader.ReadSlice(delim)
		s.ch <- true
	}()
	select {
	case <-s.ch:
		return
	case <-time.After(s.timeout):
		return nil, Timeout
	}
	return
}

// ReadLine wraps bufio.ReadLine adding timeout
func (s *TimeoutReader) ReadLine() (line []byte, isPrefix bool, err error) {
	go func() {
		line, isPrefix, err = s.Reader.ReadLine()
		s.ch <- true
	}()
	select {
	case <-s.ch:
		return
	case <-time.After(s.timeout):
		return nil, false, Timeout
	}
	return
}

// ReadBytes wraps bufio.ReadBytes adding timeout
func (s *TimeoutReader) ReadBytes(delim byte) (b []byte, err error) {
	go func() {
		b, err = s.Reader.ReadBytes(delim)
		s.ch <- true
	}()
	select {
	case <-s.ch:
		return
	case <-time.After(s.timeout):
		return nil, Timeout
	}
	return
}

// ReadString wraps bufio.ReadString adding timeout
func (s *TimeoutReader) ReadString(delim byte) (str string, err error) {
	go func() {
		str, err = s.Reader.ReadString(delim)
		s.ch <- true
	}()
	select {
	case <-s.ch:
		return
	case <-time.After(s.timeout):
		return "", Timeout
	}
	return
}
