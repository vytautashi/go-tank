package util

import (
	"log"
)

type ChannelMessage struct {
	Sender uint32
	Data   []byte
}

// Sends to channel of type `ChannelMessage`
func SendToChannelMessage(channel chan<- ChannelMessage, sender uint32, data []byte) {
	m := ChannelMessage{
		Sender: sender,
		Data:   data,
	}

	select {
	case channel <- m:
	default:
	}
}

// Sends bytes slice to channel
func SendBytesToChannel(channel chan<- []byte, data []byte) {
	select {
	case channel <- data:
	default:
	}
}

// Converts various type of data to single byte slice.
// On fail logs error.
func ToBytes(a ...interface{}) []byte {
	result := []byte{}

	for i := 0; i < len(a); i++ {
		switch v := a[i].(type) {
		case []byte:
			result = append(result, v...)
		case byte:
			result = append(result, v)
		case int16:
			result = append(result, byte(v>>8), byte(v))
		case uint16:
			result = append(result, byte(v>>8), byte(v))
		case int32:
			result = append(result, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
		case uint32:
			result = append(result, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
		default:
			log.Println("[ERROR] unknown variable type")
		}
	}

	return result
}
