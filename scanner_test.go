package bufioto

import (
	"bufio"
	"io"
	"testing"
	"time"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestTimeoutScanner_Scan(t *testing.T) {
	r, w := io.Pipe()
	scanner := NewTimeoutScanner(bufio.NewScanner(r), 500*time.Millisecond)
	stringsSend := []string{
		"foo\n", // should not time out
		"foo2\n", // should not time out
		"bar",   // should time out
	}
	go func() {
		for _, s := range stringsSend {
			w.Write([]byte(s))
		}
	}()

	timeoutCount := 0
	stringsReceive := make([]string, 0, 1)
	for scanner.Scan() {
		s := scanner.Text()
		stringsReceive = append(stringsReceive, s)
	}
	if err := scanner.Err(); err != nil {
		if err == Timeout {
			timeoutCount++
		} else {
			t.Error(err)
		}
	}

	assert.Equal(t, 1, timeoutCount, "timout count not correct")
	assert.Equal(t, 2, len(stringsReceive), "not enough messages received")
	if len(stringsReceive) > 0 {
		assert.Equal(t, strings.Replace(stringsSend[0], "\n", "", -1), stringsReceive[0])
	}
}
