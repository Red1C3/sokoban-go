package state

const (
	BLANK    = 0x0
	BOX      = 0x1
	OBSTACLE = 0x2
	PLAYER   = 0x4
	GOAL     = 0x8
)

const (
	BLANKCHAR    = " "
	BOXCHAR      = "o"
	OBSTACLECHAR = "x"
	PLAYERCHAR   = "*"
	GOALCHAR     = "."
	VSEPERATOR   = "|"
	HSEPERATOR   = "-"
)

const (
	_ = iota
	UP
	DOWN
	LEFT
	RIGHT
)

const BORDER = 1
