package server

import (
	"testing"
)

// Test: Constants
func TestConstantsServer(t *testing.T) {
	// 1: framesPerSecond
	expect := 30
	result := framesPerSecond
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 2: frameDurationInMiliseconds
	expect = 33
	result = frameDurationInMiliseconds
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 3: frameDurationInMiliseconds = 1000 / framesPerSecond
	expect = 1000 / framesPerSecond
	result = frameDurationInMiliseconds
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}
