package state

// State tiles integer representations
const (
	BLANK        = 0x0
	BOX          = 0x1
	OBSTACLE     = 0x2
	PLAYER       = 0x4
	GOAL         = 0x8
	BOXONGOAL    = GOAL | BOX
	PLAYERONGOAL = GOAL | PLAYER
)

// State tiles human-readable representation (for both input and output)
const (
	BLANKCHAR        = "⬜"
	BOXCHAR          = "🟩"
	OBSTACLECHAR     = "🟥"
	PLAYERCHAR       = "🐈"
	GOALCHAR         = "⭕"
	BOXONGOALCHAR    = "✅"
	PLAYERONGOALCHAR = "😿"
)

// Directions enum
const (
	_ = iota
	UP
	DOWN
	LEFT
	RIGHT
)

// Added border for states (should never be less than 1)
const BORDER = 1
