package client

import (
	"bytes"
	"testing"
	"time"
)

// Test: Constants
func TestConstantsMessageLength(t *testing.T) {
	// 1: msgLengthMin
	expect := 2
	result := msgLengthMin
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 2: msgLengthMax
	expect = 8
	result = msgLengthMax
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: Constants
func TestConstantsTime(t *testing.T) {
	// 1: timeoutLimit
	expect := 5 * time.Second
	result := timeoutLimit
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 2: pingingTime
	expect = 1 * time.Second
	result = pingingTime
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: Constants (timeoutLimit > pingingTime)
func TestConstantsTimeoutMoreThanPinging(t *testing.T) {
	expect := true
	result := timeoutLimit > pingingTime
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: Send()
func TestSend(t *testing.T) {
	client := New(nil, nil, nil, 1)
	go client.Send([]byte{0b0000_1100, 0b0000_0011})

	expect := []byte{0b0000_1100, 0b0000_0011}
	result := <-client.send
	if !bytes.Equal(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}
