package server

import (
	"reflect"
	"testing"
)

// Test: NewGame()
func TestNewGame(t *testing.T) {
	game := *NewGame(200, 100)
	expect := []uint16{200, 100}
	result := []uint16{game.mapWidth, game.mapheight}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: updatePlayersPositions()
func TestUpdatePlayersPositions(t *testing.T) {
	// 0: Initial setup
	game := NewGame(10, 10)
	game.players[1] = NewPlayer()
	game.players[2] = NewPlayer()
	game.players[1].updateInput(moveUp)
	game.players[2].updateInput(moveRight)

	expect := []int16{0, 0, 0, 0}
	result := []int16{
		game.players[1].x,
		game.players[1].y,
		game.players[2].x,
		game.players[2].y}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 1:
	game.updatePlayersPositions()
	expect = []int16{0, 2, 2, 0}
	result = []int16{
		game.players[1].x,
		game.players[1].y,
		game.players[2].x,
		game.players[2].y}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 2:
	game.updatePlayersPositions()
	expect = []int16{0, 4, 4, 0}
	result = []int16{
		game.players[1].x,
		game.players[1].y,
		game.players[2].x,
		game.players[2].y}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 3:
	game.players[1].updateInput(moveDown)
	game.updatePlayersPositions()
	expect = []int16{0, 2, 6, 0}
	result = []int16{
		game.players[1].x,
		game.players[1].y,
		game.players[2].x,
		game.players[2].y}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 4:
	game.updatePlayersPositions()
	expect = []int16{0, 0, 8, 0}
	result = []int16{
		game.players[1].x,
		game.players[1].y,
		game.players[2].x,
		game.players[2].y}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 5: player[1] reaches end of map in y axis, so position stay same
	game.updatePlayersPositions()
	expect = []int16{0, 0, 10, 0}
	result = []int16{
		game.players[1].x,
		game.players[1].y,
		game.players[2].x,
		game.players[2].y}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 6: player[2] reaches end of map in x axis, so position stay same
	game.updatePlayersPositions()
	expect = []int16{0, 0, 10, 0}
	result = []int16{
		game.players[1].x,
		game.players[1].y,
		game.players[2].x,
		game.players[2].y}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 7:
	game.players[1].updateInput(moveUp)
	game.players[2].updateInput(moveUp)
	game.updatePlayersPositions()
	expect = []int16{0, 2, 10, 2}
	result = []int16{
		game.players[1].x,
		game.players[1].y,
		game.players[2].x,
		game.players[2].y}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 8:
	game.updatePlayersPositions()
	game.updatePlayersPositions()
	game.updatePlayersPositions()
	game.updatePlayersPositions()
	game.updatePlayersPositions()
	expect = []int16{0, 10, 10, 10}
	result = []int16{
		game.players[1].x,
		game.players[1].y,
		game.players[2].x,
		game.players[2].y}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 9:
	game.players[1].updateInput(moveRight)
	game.players[2].updateInput(moveLeft)
	game.updatePlayersPositions()
	expect = []int16{2, 10, 8, 10}
	result = []int16{
		game.players[1].x,
		game.players[1].y,
		game.players[2].x,
		game.players[2].y}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 10:
	game.updatePlayersPositions()
	game.updatePlayersPositions()
	game.updatePlayersPositions()
	game.updatePlayersPositions()
	game.updatePlayersPositions()
	expect = []int16{10, 10, 0, 10}
	result = []int16{
		game.players[1].x,
		game.players[1].y,
		game.players[2].x,
		game.players[2].y}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: updatePlayersActions()
func TestUpdatePlayersActions(t *testing.T) {
	// 0: Initial setup
	game := NewGame(10, 10)
	game.players[1] = NewPlayer()
	game.players[1].updateInput(moveDown)

	expect := &Player{
		ymove:        -2,
		fireGun:      false,
		bulletRecoil: 0,
	}
	result := game.players[1]
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 1:
	game.updatePlayersActions()
	expect = &Player{
		ymove:        -2,
		fireGun:      false,
		bulletRecoil: 0,
	}
	result = game.players[1]
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 2:
	game.players[1].updateInput(fireGun)
	game.updatePlayersActions()
	expect = &Player{
		ymove:        -2,
		fireGun:      true,
		bulletRecoil: 1,
	}
	result = game.players[1]
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 3:
	game.updatePlayersActions()
	game.updatePlayersActions()
	game.updatePlayersActions()
	game.updatePlayersActions()
	game.updatePlayersActions()
	game.updatePlayersActions()
	game.updatePlayersActions()
	game.updatePlayersActions()
	expect = &Player{
		ymove:        -2,
		fireGun:      true,
		bulletRecoil: 9,
	}
	result = game.players[1]
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect2 := 0
	result2 := len(game.bulletManager.bullets)
	if result2 != expect2 {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}

	// 4: fire bullets
	game.updatePlayersActions()
	expect = &Player{
		ymove:        -2,
		fireGun:      false,
		bulletRecoil: 0,
	}
	result = game.players[1]
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	// Added bullet
	expect2 = 1
	result2 = len(game.bulletManager.bullets)
	if result2 != expect2 {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}
	// Test bullet data
	expect3 := Bullet{
		playerID: 1,
		lifeTime: 120,
		x:        0,
		y:        0,
		xAdd:     0,
		yAdd:     -10,
	}
	result3 := game.bulletManager.bullets[0]
	if result3 != expect3 {
		t.Fatalf(`result = %v, expect = %v`, result3, expect3)
	}

	// 5:
	game.updatePlayersActions()
	expect = &Player{
		ymove:        -2,
		fireGun:      false,
		bulletRecoil: 0,
	}
	result = game.players[1]
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect2 = 1
	result2 = len(game.bulletManager.bullets)
	if result2 != expect2 {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}

	// 5:
	game.players[1].updateInput(fireGun)
	game.updatePlayersActions()
	expect = &Player{
		ymove:        -2,
		fireGun:      true,
		bulletRecoil: 1,
	}
	result = game.players[1]
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect2 = 1
	result2 = len(game.bulletManager.bullets)
	if result2 != expect2 {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}
}
