package server

import (
	"testing"
)

// Test: nextClientID() generates `unique` id.
func TestNextClientID_return_1(t *testing.T) {
	var expect int32 = 1
	result := nextClientID()
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: nextClientID() generates `unique` id.
// Returns incremented id.
func TestNextClientID_return_2(t *testing.T) {
	var expect int32 = 2
	result := nextClientID()
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}
