package algor

import "strings"

const(
	UP="u"
	DOWN="d"
	LEFT="l"
	RIGHT="r"
)

var(
	PROMPT=strings.Join([]string{UP,DOWN,LEFT,RIGHT},", ")+ ">"
)