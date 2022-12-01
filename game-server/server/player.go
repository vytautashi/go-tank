package server

// Used for identifying client input commands based on which bits are set.
const (
	moveRight = 0b0000_0001
	moveLeft  = 0b0000_0010
	moveUp    = 0b0000_0100
	moveDown  = 0b0000_1000
	fireGun   = 0b0001_0000
)

type Player struct {
	x            int16
	y            int16
	xmove        int16
	ymove        int16
	fireGun      bool
	bulletRecoil int
}

// Creates new player.
func NewPlayer() *Player {
	return &Player{
		x:            0,
		y:            0,
		xmove:        0,
		ymove:        2,
		fireGun:      false,
		bulletRecoil: 0,
	}
}

// Updates player inputs based on `cmd` using bitwise AND(&) operator.
func (p *Player) updateInput(cmd byte) {
	if cmd&moveRight == moveRight {
		p.xmove = 2
		p.ymove = 0
	} else if cmd&moveLeft == moveLeft {
		p.xmove = -2
		p.ymove = 0
	} else if cmd&moveUp == moveUp {
		p.xmove = 0
		p.ymove = 2
	} else if cmd&moveDown == moveDown {
		p.xmove = 0
		p.ymove = -2
	}

	if cmd&fireGun == fireGun {
		p.fireGun = true
	}
}
