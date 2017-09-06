package bufioto

import (
	"bufio"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeoutReader_ReadString(t *testing.T) {
	r, w := io.Pipe()
	reader := NewTimeoutReader(*bufio.NewReader(r), 500*time.Millisecond)
	stringsSend := []string{
		"foo\n", // should not time out
		"bar",   // should time out
	}
	go func() {
		for _, s := range stringsSend {
			w.Write([]byte(s))
		}
	}()

	timeoutCount := 0
	stringsReceive := make([]string, 0, 1)
	for i := 0; i < len(stringsSend); i++ {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == Timeout {
				timeoutCount++
				fmt.Println(err)
				continue
			}
			t.Error(err)
			break
		}
		fmt.Println("got:", str)
		stringsReceive = append(stringsReceive, str)
	}

	assert.Equal(t, 1, timeoutCount, "timout count not correct")
	assert.Equal(t, 1, len(stringsReceive), "not enough messages received")
	if len(stringsReceive) > 0 {
		assert.Equal(t, stringsSend[0], stringsReceive[0])
	}
}
