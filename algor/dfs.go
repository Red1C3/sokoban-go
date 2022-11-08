package algor

import (
	"container/list"
	"sokoban/state"
)

type Dfs struct {
}

func (d *Dfs) Steps() int {
	return 0
}

func (d *Dfs) Search(start state.State) *state.State {
	visited := make([]state.State, 0)

	stack := list.New()
	stack.PushBack(start)

	for stack.Len() > 0 {
		top := stack.Remove(stack.Back()).(state.State)

		if top.IsSolved() {
			return &top
		}

		if !isVisited(visited, top) {
			children := top.States()
			for _, v := range children {
				stack.PushBack(v)
			}
			visited=append(visited,top)
		}
	}

	return nil
}