package algor

import "sokoban/state"

type Algor interface{
	Steps()int
	Search(start state.State)*state.State
}

func isVisited(states []state.State, state state.State) bool {
	for _, v := range states {
		if v.Equals(&state) {
			return true
		}
	}
	return false
}
