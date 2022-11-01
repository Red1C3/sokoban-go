package algor

import "sokoban/state"

type Algor interface{
	Steps()int
	Search(start state.State)state.State
}