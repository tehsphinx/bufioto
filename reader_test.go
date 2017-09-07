package bufioto

import (
	"bufio"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeoutReader_ReadString(t *testing.T) {
	r, w := io.Pipe()
	reader := NewTimeoutReader(bufio.NewReader(r), 500*time.Millisecond)
	stringsSend := []string{
		"foo\n",  // should not time out
		"foo2\n", // should not time out
		"bar",    // should time out
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
				break
			}
			t.Error(err)
			break
		}
		stringsReceive = append(stringsReceive, str)
	}

	assert.Equal(t, 1, timeoutCount, "timout count not correct")
	assert.Equal(t, 2, len(stringsReceive), "not enough messages received")
	if len(stringsReceive) > 0 {
		assert.Equal(t, stringsSend[0], stringsReceive[0])
	}
}

//Failing: After timeout blocking read call is still waiting and will swallow some input
//func TestTimeoutReader_ReadString_AfterTimeout(t *testing.T) {
//	r, w := io.Pipe()
//	reader := NewTimeoutReader(bufio.NewReader(r), 500*time.Millisecond)
//	stringsSend := []string{
//		"foo\n", // should not time out
//		"bar",   // should time out
//	}
//	go func() {
//		for _, s := range stringsSend {
//			w.Write([]byte(s))
//		}
//	}()
//
//	timeoutCount := 0
//	stringsReceive := make([]string, 0, 1)
//	for i := 0; i < len(stringsSend); i++ {
//		str, err := reader.ReadString('\n')
//		if err != nil {
//			if err == Timeout {
//				timeoutCount++
//				fmt.Println(err)
//				continue
//			}
//			t.Error(err)
//			break
//		}
//		fmt.Println("got:", str)
//		stringsReceive = append(stringsReceive, str)
//	}
//
//	assert.Equal(t, 1, timeoutCount, "timout count not correct")
//	assert.Equal(t, 1, len(stringsReceive), "not enough messages received")
//	if len(stringsReceive) > 0 {
//		assert.Equal(t, stringsSend[0], stringsReceive[0])
//	}
//
//	// test the same again to see if continuation works or if previous read swallows information
//	go func() {
//		for _, s := range stringsSend {
//			w.Write([]byte(s))
//		}
//	}()
//
//	timeoutCount = 0
//	stringsReceive = make([]string, 0, 1)
//	for i := 0; i < len(stringsSend); i++ {
//		str, err := reader.ReadString('\n')
//		if err != nil {
//			if err == Timeout {
//				timeoutCount++
//				fmt.Println(err)
//				continue
//			}
//			t.Error(err)
//			break
//		}
//		fmt.Println("got:", str)
//		stringsReceive = append(stringsReceive, str)
//	}
//
//	assert.Equal(t, 1, timeoutCount, "timout count not correct")
//	assert.Equal(t, 1, len(stringsReceive), "not enough messages received")
//	if len(stringsReceive) > 0 {
//		assert.Equal(t, stringsSend[0], stringsReceive[0])
//	}
//}
