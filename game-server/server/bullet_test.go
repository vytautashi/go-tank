package server

import (
	"reflect"
	"testing"
)

// Test: Constants
func TestConstantsBullet(t *testing.T) {
	// 1: bulletExploded
	expect := -1
	result := bulletExploded
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 2: bulletSizeInBytes
	expect = 4
	result = bulletSizeInBytes
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: bulletSizeInBytes
func TestBulletSizeInBytes_equals_to_bullet_x_plus_y_type_size(t *testing.T) {
	var bullet Bullet

	expect := bulletSizeInBytes
	result := int(reflect.TypeOf(bullet.x).Size() + reflect.TypeOf(bullet.y).Size())
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: newBullet()
func TestNewBullet(t *testing.T) {
	// 1:
	expect := Bullet{
		playerID: 22,
		lifeTime: 120,
		x:        0,
		y:        0,
		xAdd:     0,
		yAdd:     10,
	}
	result := newBullet(22, &Player{x: 0, y: 0, xmove: 0, ymove: 2})
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 2:
	expect = Bullet{
		playerID: 6,
		lifeTime: 120,
		x:        10,
		y:        12,
		xAdd:     -15,
		yAdd:     0,
	}
	result = newBullet(6, &Player{x: 10, y: 12, xmove: -3, ymove: 0})
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: addBullet()
func TestAddBullet(t *testing.T) {
	// 0: Initial setup
	bulletManager := newBulletManager()
	player := Player{x: 0, y: 0, xmove: 0, ymove: 2}
	expect := 0
	result := len(bulletManager.bullets)
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 1:
	bulletManager.addBullet(1, &player)
	expect = 1
	result = len(bulletManager.bullets)
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 2:
	bulletManager.addBullet(1, &player)
	bulletManager.addBullet(1, &player)
	bulletManager.addBullet(2, &player)
	bulletManager.addBullet(2, &player)
	expect = 5
	result = len(bulletManager.bullets)
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: updateBulletsMovement()
func TestUpdateBulletsMovement(t *testing.T) {
	// 0: Initial setup
	bulletManager := newBulletManager()
	bulletManager.bullets = append(bulletManager.bullets,
		Bullet{lifeTime: 120, x: 8, y: 16, xAdd: 10, yAdd: 0},
		Bullet{lifeTime: 80, x: 0, y: 0, xAdd: 0, yAdd: 5},
	)
	expect := 2
	result := len(bulletManager.bullets)
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 1:
	bulletManager.updateBulletsMovement()
	expect2 := []Bullet{
		{lifeTime: 119, x: 18, y: 16, xAdd: 10, yAdd: 0},
		{lifeTime: 79, x: 0, y: 5, xAdd: 0, yAdd: 5},
	}
	result2 := bulletManager.bullets
	if !reflect.DeepEqual(result2, expect2) {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}

	// 2:
	bulletManager.updateBulletsMovement()
	bulletManager.updateBulletsMovement()
	bulletManager.updateBulletsMovement()
	expect2 = []Bullet{
		{lifeTime: 116, x: 48, y: 16, xAdd: 10, yAdd: 0},
		{lifeTime: 76, x: 0, y: 20, xAdd: 0, yAdd: 5},
	}
	result2 = bulletManager.bullets
	if !reflect.DeepEqual(result2, expect2) {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}
}

// Test: updateBulletsCollision()
func TestUpdateBulletsCollision(t *testing.T) {
	bulletManager := newBulletManager()
	players := make(map[uint32]*Player)
	players[1] = &Player{x: 50, y: 100}
	players[2] = &Player{x: 200, y: 300}

	bulletManager.bullets = append(bulletManager.bullets,
		Bullet{lifeTime: 100, x: 29, y: 100},
		Bullet{lifeTime: 100, x: 71, y: 100},
		Bullet{lifeTime: 100, x: 50, y: 79},
		Bullet{lifeTime: 100, x: 50, y: 121},
	)

	// 1: No collisions
	bulletManager.updateBulletsCollision(players)
	expect := []Bullet{
		{lifeTime: 100, x: 29, y: 100},
		{lifeTime: 100, x: 71, y: 100},
		{lifeTime: 100, x: 50, y: 79},
		{lifeTime: 100, x: 50, y: 121},
	}
	result := bulletManager.bullets
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect2 := map[uint32]*Player{
		1: {x: 50, y: 100},
		2: {x: 200, y: 300},
	}
	result2 := players
	if !reflect.DeepEqual(result2, expect2) {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}

	// 2: Player collides x left bullet
	players[1] = &Player{x: 49, y: 100}
	bulletManager.updateBulletsCollision(players)
	expect = []Bullet{
		{lifeTime: -1, x: 29, y: 100},
		{lifeTime: 100, x: 71, y: 100},
		{lifeTime: 100, x: 50, y: 79},
		{lifeTime: 100, x: 50, y: 121},
	}
	result = bulletManager.bullets
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect2 = map[uint32]*Player{
		1: {x: 0, y: 0, ymove: 2},
		2: {x: 200, y: 300},
	}
	result2 = players
	if !reflect.DeepEqual(result2, expect2) {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}

	// 3: Player collides x right bullet
	players[1] = &Player{x: 51, y: 100}
	bulletManager.updateBulletsCollision(players)
	expect = []Bullet{
		{lifeTime: -1, x: 29, y: 100},
		{lifeTime: -1, x: 71, y: 100},
		{lifeTime: 100, x: 50, y: 79},
		{lifeTime: 100, x: 50, y: 121},
	}
	result = bulletManager.bullets
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect2 = map[uint32]*Player{
		1: {x: 0, y: 0, ymove: 2},
		2: {x: 200, y: 300},
	}
	result2 = players
	if !reflect.DeepEqual(result2, expect2) {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}

	// 4: Player collides y bottom bullet
	players[1] = &Player{x: 50, y: 99}
	bulletManager.updateBulletsCollision(players)
	expect = []Bullet{
		{lifeTime: -1, x: 29, y: 100},
		{lifeTime: -1, x: 71, y: 100},
		{lifeTime: -1, x: 50, y: 79},
		{lifeTime: 100, x: 50, y: 121},
	}
	result = bulletManager.bullets
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect2 = map[uint32]*Player{
		1: {x: 0, y: 0, ymove: 2},
		2: {x: 200, y: 300},
	}
	result2 = players
	if !reflect.DeepEqual(result2, expect2) {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}

	// 5: Player collides y top bullet
	players[1] = &Player{x: 50, y: 101}
	bulletManager.updateBulletsCollision(players)
	expect = []Bullet{
		{lifeTime: -1, x: 29, y: 100},
		{lifeTime: -1, x: 71, y: 100},
		{lifeTime: -1, x: 50, y: 79},
		{lifeTime: -1, x: 50, y: 121},
	}
	result = bulletManager.bullets
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect2 = map[uint32]*Player{
		1: {x: 0, y: 0, ymove: 2},
		2: {x: 200, y: 300},
	}
	result2 = players
	if !reflect.DeepEqual(result2, expect2) {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}

	// 6: Same bullet must not collide/explode twice
	bulletManager.bullets[0] = Bullet{lifeTime: 100, x: 29, y: 100}
	players[1] = &Player{x: 30, y: 100}
	players[2] = &Player{x: 30, y: 100}
	bulletManager.updateBulletsCollision(players)
	expect = []Bullet{
		{lifeTime: -1, x: 29, y: 100},
		{lifeTime: -1, x: 71, y: 100},
		{lifeTime: -1, x: 50, y: 79},
		{lifeTime: -1, x: 50, y: 121},
	}
	result = bulletManager.bullets
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect2 = map[uint32]*Player{
		1: {x: 0, y: 0, ymove: 2},
		2: {x: 30, y: 100},
	}
	result2 = players
	if !reflect.DeepEqual(result2, expect2) {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}

	// 7: bullet do not collide with its owner
	bulletManager.bullets[0] = Bullet{playerID: 1, lifeTime: 100, x: 29, y: 100}
	players[1] = &Player{x: 30, y: 100}
	players[2] = &Player{x: 200, y: 300}
	bulletManager.updateBulletsCollision(players)
	expect = []Bullet{
		{playerID: 1, lifeTime: 100, x: 29, y: 100},
		{lifeTime: -1, x: 71, y: 100},
		{lifeTime: -1, x: 50, y: 79},
		{lifeTime: -1, x: 50, y: 121},
	}
	result = bulletManager.bullets
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect2 = map[uint32]*Player{
		1: {x: 30, y: 100},
		2: {x: 200, y: 300},
	}
	result2 = players
	if !reflect.DeepEqual(result2, expect2) {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}

	// 8: Many(2) collisions
	bulletManager.bullets[0] = Bullet{lifeTime: 100, x: 30, y: 100}
	bulletManager.bullets[1] = Bullet{lifeTime: 100, x: 30, y: 100}
	bulletManager.bullets[2] = Bullet{lifeTime: 100, x: 1000, y: 1000}
	bulletManager.bullets[3] = Bullet{lifeTime: 100, x: 200, y: 300}
	players[1] = &Player{x: 30, y: 100}
	players[2] = &Player{x: 200, y: 300}
	bulletManager.updateBulletsCollision(players)
	expect = []Bullet{
		{lifeTime: -1, x: 30, y: 100},
		{lifeTime: 100, x: 30, y: 100},
		{lifeTime: 100, x: 1000, y: 1000},
		{lifeTime: -1, x: 200, y: 300},
	}
	result = bulletManager.bullets
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect2 = map[uint32]*Player{
		1: {x: 0, y: 0, ymove: 2},
		2: {x: 0, y: 0, ymove: 2},
	}
	result2 = players
	if !reflect.DeepEqual(result2, expect2) {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}
}

// Test: removeBullets()
func TestRemoveBullets(t *testing.T) {
	// 0: Initial setup
	bulletManager := newBulletManager()
	bulletManager.bullets = append(bulletManager.bullets,
		Bullet{lifeTime: 120, x: 8, y: 16, xAdd: 10, yAdd: 0},
		Bullet{lifeTime: 80, x: 0, y: 0, xAdd: 0, yAdd: 5},
		Bullet{lifeTime: 0, x: 0, y: 0, xAdd: 0, yAdd: 5},
		Bullet{lifeTime: -1, x: 0, y: 0, xAdd: 0, yAdd: 5},
		Bullet{lifeTime: 80, x: 0, y: 0, xAdd: 0, yAdd: 5},
	)
	expect := 5
	result := len(bulletManager.bullets)
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 1:
	bulletManager.removeBullets()
	expect2 := []Bullet{
		{lifeTime: 120, x: 8, y: 16, xAdd: 10, yAdd: 0},
		{lifeTime: 80, x: 0, y: 0, xAdd: 0, yAdd: 5},
		{lifeTime: 80, x: 0, y: 0, xAdd: 0, yAdd: 5},
	}
	result2 := bulletManager.bullets
	if !reflect.DeepEqual(result2, expect2) {
		t.Fatalf(`result = %v, expect = %v`, result2, expect2)
	}
}

// Test: getBulletsData()
func TestGetBulletsData(t *testing.T) {
	bulletManager := newBulletManager()
	bulletManager.bullets = append(bulletManager.bullets,
		Bullet{x: 8, y: 16},
		Bullet{x: 0, y: 33},
	)

	expect := []byte{
		0b0000_0000, 0b0000_1000, // 8
		0b0000_0000, 0b0001_0000, // 16
		0b0000_0000, 0b0000_0000, // 0
		0b0000_0000, 0b0010_0001, // 33
	}
	result := bulletManager.getBulletsData()
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}
