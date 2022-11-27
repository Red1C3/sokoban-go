package state

const (
	BLANK        = 0x0
	BOX          = 0x1
	OBSTACLE     = 0x2
	PLAYER       = 0x4
	GOAL         = 0x8
	BOXONGOAL    = GOAL | BOX
	PLAYERONGOAL = GOAL | PLAYER
)

const (
	BLANKCHAR        = "⬜"
	BOXCHAR          = "🟩"
	OBSTACLECHAR     = "🟥"
	PLAYERCHAR       = "🐈"
	GOALCHAR         = "⭕"
	BOXONGOALCHAR    = "✅"
	PLAYERONGOALCHAR = "😿"
)

const (
	_ = iota
	UP
	DOWN
	LEFT
	RIGHT
)

const BORDER = 1
