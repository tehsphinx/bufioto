package bufioto

import (
	"bufio"
	"errors"
	"io"
	"time"
)

var Timeout = errors.New("timeout on reading")

// TimeoutReader is a wrapper for a bufio reader adding timeout
type TimeoutReader struct {
	*bufio.Reader
	timeout time.Duration
}

// NewTimeoutReader creates a new TimeoutReader
func NewTimeoutReader(reader io.Reader, timeout time.Duration) *TimeoutReader {
	return &TimeoutReader{
		Reader:  bufio.NewReader(reader),
		timeout: timeout,
	}
}

// Peek wraps bufio.Reader.Peek adding timeout
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

	// no timout: wait for reader result
	if s.timeout == 0 {
		<-ch
		return
	}

	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return nil, Timeout
	}
}

// Discard wraps bufio.Reader.Discard adding timeout
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

	// no timout: wait for reader result
	if s.timeout == 0 {
		<-ch
		return
	}

	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return 0, Timeout
	}
}

// Read wraps bufio.Reader.Read adding timeout
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

	// no timout: wait for reader result
	if s.timeout == 0 {
		<-ch
		return
	}

	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return 0, Timeout
	}
}

// ReadByte wraps bufio.Reader.ReadByte adding timeout
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

	// no timout: wait for reader result
	if s.timeout == 0 {
		<-ch
		return
	}

	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return 0, Timeout
	}
}

// ReadRune wraps bufio.Reader.ReadRune adding timeout
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

	// no timout: wait for reader result
	if s.timeout == 0 {
		<-ch
		return
	}

	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return 0, 0, Timeout
	}
}

// ReadSlice wraps bufio.Reader.ReadSlice adding timeout
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

	// no timout: wait for reader result
	if s.timeout == 0 {
		<-ch
		return
	}

	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return nil, Timeout
	}
}

// ReadLine wraps bufio.Reader.ReadLine adding timeout
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

	// no timout: wait for reader result
	if s.timeout == 0 {
		<-ch
		return
	}

	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return nil, false, Timeout
	}
}

// ReadBytes wraps bufio.Reader.ReadBytes adding timeout
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

	// no timout: wait for reader result
	if s.timeout == 0 {
		<-ch
		return
	}

	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return nil, Timeout
	}
}

// ReadString wraps bufio.Reader.ReadString adding timeout
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

	// no timout: wait for reader result
	if s.timeout == 0 {
		<-ch
		return
	}

	select {
	case <-ch:
		return
	case <-time.After(s.timeout):
		chTo <- true
		return "", Timeout
	}
}
