package algor

import "strings"

// Keyboard input keys for algor.NoAlgor
const (
	UP    = "u"
	DOWN  = "d"
	LEFT  = "l"
	RIGHT = "r"
)

// On screen prompt when using algor.NoAlgor
var (
	PROMPT = strings.Join([]string{UP, DOWN, LEFT, RIGHT}, ", ") + ">"
)
