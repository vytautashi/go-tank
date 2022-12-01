package util

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

// Test: SendToChannelMessage()
func TestSendToChannelMessage(t *testing.T) {
	channel := make(chan ChannelMessage)
	go SendToChannelMessage(channel, 33, []byte{0b0000_1100, 0b0000_1111})

	result := <-channel
	expect := ChannelMessage{
		Sender: 33,
		Data:   []byte{0b0000_1100, 0b0000_1111},
	}

	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: SendBytesToChannel()
func TestSendBytesToChannel(t *testing.T) {
	channel := make(chan []byte)
	go SendBytesToChannel(channel, []byte{0b0000_1100, 0b0000_1111})

	result := <-channel
	expect := []byte{0b0000_1100, 0b0000_1111}

	if !bytes.Equal(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: ToBytes()
// Converts single byte to byte slice
func TestToBytes_byte(t *testing.T) {
	expect := []byte{0b0000_1100}
	result := ToBytes(byte(0b0000_1100))
	if !bytes.Equal(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: ToBytes()
func TestToBytes_byte_byte(t *testing.T) {
	expect := []byte{0b0000_1100, 0b1111_0000}
	result := ToBytes(byte(0b0000_1100), byte(0b1111_0000))
	if !bytes.Equal(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: ToBytes()
func TestToBytes_byte_byte_slice(t *testing.T) {
	expect := []byte{0b0000_1100, 0b0000_0001, 0b0001_0000}
	result := ToBytes(byte(0b0000_1100), []byte{0b0000_0001, 0b0001_0000})
	if !bytes.Equal(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: ToBytes()
func TestToBytes_byte_int16(t *testing.T) {
	expect := []byte{0b0000_1100, 0b0000_0010, 0b0000_0001}
	result := ToBytes(byte(0b0000_1100), int16(0b0000_0010_0000_0001))
	if !bytes.Equal(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: ToBytes()
func TestToBytes_multiple(t *testing.T) {
	expect := []byte{
		0b0000_1100,
		0b0000_0010, 0b0000_0001,
		0b1000_0011, 0b1000_1111,
		0b0001_0000, 0b0000_1000, 0b0000_0100, 0b0000_0010,
		0b1111_1111, 0b0011_1111, 0b1111_1011, 0b1100_1111}
	result := ToBytes(
		byte(0b0000_1100),
		int16(0b0000_0010_0000_0001),
		uint16(0b1000_0011_1000_1111),
		int32(0b0001_0000_0000_1000_0000_0100_0000_0010),
		uint32(0b1111_1111_0011_1111_1111_1011_1100_1111))
	if !bytes.Equal(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: ToBytes()
func TestToBytes_returns_empty_byte_slice_when_unknow_type(t *testing.T) {
	expect := []byte{}
	result := ToBytes(true)
	if !bytes.Equal(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	expect = []byte{}
	result = ToBytes("hello")
	if !bytes.Equal(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: ToBytes()
func TestToBytes_logs_error_when_unknow_type(t *testing.T) {
	// 0: initial setup for testing log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	// 1: known type
	ToBytes(byte(0b0000_1100))
	expect := ""
	result := buf.String()
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 2: unknown type, logs error
	ToBytes(true)
	expect = "[ERROR] unknown variable type"
	result = buf.String()
	if !strings.Contains(result, expect) {
		t.Fatalf(`result = "%v", expect(to contain) = "%v"`, result, expect)
	}
}
