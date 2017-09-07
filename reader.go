package bufioto

import (
	"bufio"
	"errors"
	"time"
)

var Timeout = errors.New("Timeout on reading")

// TimeoutReader is a wrapper for a bufio reader adding timeout
type TimeoutReader struct {
	*bufio.Reader
	timeout time.Duration
}

// NewTimeoutReader creates a new TimeoutReader
func NewTimeoutReader(reader *bufio.Reader, timeout time.Duration) *TimeoutReader {
	return &TimeoutReader{
		Reader:  reader,
		timeout: timeout,
	}
}

// Peek wraps bufio.Peek adding timeout
func (s *TimeoutReader) Peek(n int) (b []byte, err error) {
	ch := make(chan bool)
	chTo := make(chan bool, 1)
	go func() {
		b, err = s.Reader.Peek(n)
		select {
		case ch <- true:
		case <-chTo:
		}
	}()
	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return nil, Timeout
	}
	return
}

// Discard wraps bufio.Discard adding timeout
func (s *TimeoutReader) Discard(n int) (discarded int, err error) {
	ch := make(chan bool)
	chTo := make(chan bool, 1)
	go func() {
		discarded, err = s.Reader.Discard(n)
		select {
		case ch <- true:
		case <-chTo:
		}
	}()
	//s.Reader.Reset()
	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return 0, Timeout
	}
	return
}

// Read wraps bufio.Read adding timeout
func (s *TimeoutReader) Read(buf []byte) (n int, err error) {
	ch := make(chan bool)
	chTo := make(chan bool, 1)
	go func() {
		n, err = s.Reader.Read(buf)
		select {
		case ch <- true:
		case <-chTo:
		}
	}()
	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return 0, Timeout
	}
	return
}

// ReadByte wraps bufio.ReadByte adding timeout
func (s *TimeoutReader) ReadByte() (b byte, err error) {
	ch := make(chan bool)
	chTo := make(chan bool, 1)
	go func() {
		b, err = s.Reader.ReadByte()
		select {
		case ch <- true:
		case <-chTo:
		}
	}()
	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return 0, Timeout
	}
	return
}

// ReadRune wraps bufio.ReadRune adding timeout
func (s *TimeoutReader) ReadRune() (r rune, size int, err error) {
	ch := make(chan bool)
	chTo := make(chan bool, 1)
	go func() {
		r, size, err = s.Reader.ReadRune()
		select {
		case ch <- true:
		case <-chTo:
		}
	}()
	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return 0, 0, Timeout
	}
	return
}

// ReadSlice wraps bufio.ReadSlice adding timeout
func (s *TimeoutReader) ReadSlice(delim byte) (line []byte, err error) {
	ch := make(chan bool)
	chTo := make(chan bool, 1)
	go func() {
		line, err = s.Reader.ReadSlice(delim)
		select {
		case ch <- true:
		case <-chTo:
		}
	}()
	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return nil, Timeout
	}
	return
}

// ReadLine wraps bufio.ReadLine adding timeout
func (s *TimeoutReader) ReadLine() (line []byte, isPrefix bool, err error) {
	ch := make(chan bool)
	chTo := make(chan bool, 1)
	go func() {
		line, isPrefix, err = s.Reader.ReadLine()
		select {
		case ch <- true:
		case <-chTo:
		}
	}()
	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return nil, false, Timeout
	}
	return
}

// ReadBytes wraps bufio.ReadBytes adding timeout
func (s *TimeoutReader) ReadBytes(delim byte) (b []byte, err error) {
	ch := make(chan bool)
	chTo := make(chan bool, 1)
	go func() {
		b, err = s.Reader.ReadBytes(delim)
		select {
		case ch <- true:
		case <-chTo:
		}
	}()
	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return nil, Timeout
	}
	return
}

// ReadString wraps bufio.ReadString adding timeout
func (s *TimeoutReader) ReadString(delim byte) (str string, err error) {
	ch := make(chan bool)
	chTo := make(chan bool, 1)
	go func() {
		str, err = s.Reader.ReadString(delim)
		select {
		case ch <- true:
		case <-chTo:
		}
	}()
	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return "", Timeout
	}
	return
}
