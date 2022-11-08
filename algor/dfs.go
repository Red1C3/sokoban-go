package algor

import (
	"sokoban/state"
)

type Dfs struct {
}

func (d *Dfs) Steps() int {
	return 0
}

func (d *Dfs) Search(start state.State) state.State {
	visited := make([]state.State, 0)
	stack := make([]state.State, 0)
	stack = append(stack, start)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack=stack[:len(stack)-1]

		if top.IsSolved() {
			return top
		}

		if !isVisited(visited, top) {
			children := top.States()
			for _, v := range children {
				stack=append(stack,v)
			}
			visited=append(visited,top)
		}
	}
	return state.State{}
}

func isVisited(states []state.State, state state.State) bool {
	for _, v := range states {
		if v.Equals(&state) {
			return true
		}
	}
	return false
}
