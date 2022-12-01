package server

import (
	"testing"
)

// Test: Constants (codes of commands that are sent to clients)
func TestConstantsCommandsToClient(t *testing.T) {
	// 1: InitPlayer
	var expect uint8 = 0
	result := CommandToClientInitPlayer
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 2: InitMap
	expect = 1
	result = CommandToClientInitMap
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 3: PlayersPositions
	expect = 2
	result = CommandToClientPlayersPositions
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 4: BulletsPositions
	expect = 3
	result = CommandToClientBulletsPositions
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: Constants (codes of commands that are received from clients)
func TestConstantsCommandFromClient(t *testing.T) {
	// 1: UpdateInput
	var expect uint8 = 100
	result := CommandFromClientUpdateInput
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}
