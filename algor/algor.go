// Package algor provides multiple algorithm that implements algor interface
package algor

import "sokoban/state"

// Algor is an interface for sokoban solving algorithms
// Search method is the core of the interface, it takes
// the start state as start and returns the final state.
// Steps returns the number of steps the algorithms had to
// take to get to the returned state.
type Algor interface {
	Steps() int
	Search(start state.State) *state.State
}

func isVisited(states []state.State, state state.State) bool {
	for _, v := range states {
		if v.Equals(&state) {
			return true
		}
	}
	return false
}
