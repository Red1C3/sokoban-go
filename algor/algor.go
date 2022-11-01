package algor

import "sokoban/state"

type algor interface{
	Steps()int
	Search(start state.State)state.State
}