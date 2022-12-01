package server

type Game struct {
	players       map[uint32]*Player
	bulletManager *BulletManager
	mapWidth      uint16
	mapheight     uint16
}

func NewGame(mapWidth uint16, mapheight uint16) *Game {
	return &Game{
		players:       make(map[uint32]*Player, 600),
		bulletManager: newBulletManager(),
		mapWidth:      mapWidth,
		mapheight:     mapheight,
	}
}

func (g *Game) update() {
	g.updatePlayersPositions()
	g.updatePlayersActions()
	g.bulletManager.update(g.players)
}

func (g *Game) updatePlayersPositions() {
	for _, p := range g.players {
		p.x += p.xmove
		p.y += p.ymove

		if p.x < 0 {
			p.x = 0
		} else if p.x > int16(g.mapWidth) {
			p.x = int16(g.mapWidth)
		}

		if p.y < 0 {
			p.y = 0
		} else if p.y > int16(g.mapheight) {
			p.y = int16(g.mapheight)
		}
	}
}

func (g *Game) updatePlayersActions() {
	for id, p := range g.players {
		if p.fireGun {
			p.bulletRecoil++
			if p.bulletRecoil >= 15 {
				p.fireGun = false
				p.bulletRecoil = 0
				g.bulletManager.addBullet(id, p)
			}
		}
	}
}
