package server

import (
	"testing"
)

// Test: NewPlayer()
func TestNewPlayer(t *testing.T) {
	expect := Player{
		x:            0,
		y:            0,
		xmove:        0,
		ymove:        2,
		fireGun:      false,
		bulletRecoil: 0,
	}
	result := *NewPlayer()
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: updateInput()
func TestUpdateInput(t *testing.T) {
	player := NewPlayer()

	// 1: Move right
	player.updateInput(moveRight)
	expect := Player{
		x:            0,
		y:            0,
		xmove:        2,
		ymove:        0,
		fireGun:      false,
		bulletRecoil: 0,
	}
	result := *player
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 2: Move left
	player.updateInput(moveLeft)
	expect = Player{
		x:            0,
		y:            0,
		xmove:        -2,
		ymove:        0,
		fireGun:      false,
		bulletRecoil: 0,
	}
	result = *player
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 3: Move up
	player.updateInput(moveUp)
	expect = Player{
		x:            0,
		y:            0,
		xmove:        0,
		ymove:        2,
		fireGun:      false,
		bulletRecoil: 0,
	}
	result = *player
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 4: Move down
	player.updateInput(moveDown)
	expect = Player{
		x:            0,
		y:            0,
		xmove:        0,
		ymove:        -2,
		fireGun:      false,
		bulletRecoil: 0,
	}
	result = *player
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 5: Fire gun
	player.updateInput(fireGun)
	expect = Player{
		x:            0,
		y:            0,
		xmove:        0,
		ymove:        -2,
		fireGun:      true,
		bulletRecoil: 0,
	}
	result = *player
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 6: Fire gun, MUST NOT toggle `fireGun` to false
	player.updateInput(fireGun)
	expect = Player{
		x:            0,
		y:            0,
		xmove:        0,
		ymove:        -2,
		fireGun:      true,
		bulletRecoil: 0,
	}
	result = *player
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 7: Move down, MUST NOT toggle `ymove` to false
	player.updateInput(moveDown)
	expect = Player{
		x:            0,
		y:            0,
		xmove:        0,
		ymove:        -2,
		fireGun:      true,
		bulletRecoil: 0,
	}
	result = *player
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}
