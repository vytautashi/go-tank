package server

import "github.com/vytautashi/go-tank/game-server/util"

const (
	// State of bullet
	bulletExploded = -1

	// position x (2 bytes) + y (2 bytes) = 4 bytes
	bulletSizeInBytes = 4
)

type Bullet struct {
	// Owner of bullet
	playerID uint32
	lifeTime int16
	x        int16
	y        int16

	// Bullet speed in x direction (positive or negative)
	xAdd int16

	// Bullet speed in y direction (positive or negative)
	yAdd int16
}

func newBullet(playerID uint32, player *Player) Bullet {
	return Bullet{
		playerID: playerID,
		lifeTime: 120, // 120frames = 4 seconds
		x:        player.x,
		y:        player.y,
		xAdd:     player.xmove * 5,
		yAdd:     player.ymove * 5,
	}
}

type BulletManager struct {
	bullets []Bullet
}

// Creates bullet manager.
func newBulletManager() *BulletManager {
	return &BulletManager{
		bullets: make([]Bullet, 0, 1024),
	}
}

// Adds new bullet to bullet manager.
func (bm *BulletManager) addBullet(playerID uint32, player *Player) {
	bullet := newBullet(playerID, player)
	bm.bullets = append(bm.bullets, bullet)
}

// Main BulletManager function, must be called once per frame.
func (bm *BulletManager) update(players map[uint32]*Player) {
	bm.updateBulletsMovement()
	bm.updateBulletsCollision(players)
	bm.removeBullets()
}

// Update bullet position/movement, lifeTime.
func (bm *BulletManager) updateBulletsMovement() {
	for i := range bm.bullets {
		bm.bullets[i].x += bm.bullets[i].xAdd
		bm.bullets[i].y += bm.bullets[i].yAdd

		bm.bullets[i].lifeTime--
	}
}

// Checks and updates bullets collision with players
func (bm *BulletManager) updateBulletsCollision(players map[uint32]*Player) {
	for id, p := range players {
		// Bounding box of player
		xLeft := p.x - 20
		xRight := p.x + 20
		yBottom := p.y - 20
		yTop := p.y + 20

		for i, bullet := range bm.bullets {
			// Checks if bullet have been exploded before
			if bullet.lifeTime == bulletExploded {
				continue
			}

			// Checks if player is owner of bullet
			if bullet.playerID == id {
				continue
			}

			// Checks if bullet collides with player
			if bullet.x >= xLeft &&
				bullet.x <= xRight &&
				bullet.y >= yBottom &&
				bullet.y <= yTop {

				bm.bullets[i].lifeTime = bulletExploded
				players[id] = NewPlayer()
				break
			}
		}
	}
}

// Removes bullets that explode(lifetime=-1) or lifetime ended
func (bm *BulletManager) removeBullets() {
	i := 0
	for _, bullet := range bm.bullets {
		if bullet.lifeTime <= 0 {
			continue
		}
		bm.bullets[i] = bullet
		i++
	}
	bm.bullets = bm.bullets[:i]
}

// Gets all bullets positions (x, y).
// Data used for sending to clients for displaying bullets in frontend.
func (bm *BulletManager) getBulletsData() []byte {
	capacity := len(bm.bullets) * bulletSizeInBytes
	bulletsData := make([]byte, 0, capacity)

	for _, b := range bm.bullets {
		bullet := util.ToBytes(b.x, b.y)
		bulletsData = append(bulletsData, bullet...)
	}
	return bulletsData
}
