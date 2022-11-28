package algor

import (
	"container/list"

	"codeberg.org/RedDeadAlice/sokoban-go/state"
)

// Dfs implements the DFS algorithm.
type Dfs struct {
	steps int
}

func (d *Dfs) Steps() int {
	return d.steps
}

func (d *Dfs) Search(start state.State) *state.State {
	visited := make([]state.State, 0)

	stack := list.New()
	stack.PushBack(start)

	for stack.Len() > 0 {
		d.steps++
		top := stack.Remove(stack.Back()).(state.State)

		if top.IsSolved() {
			return &top
		}

		if !isVisited(visited, top) {
			children := top.States()
			for _, v := range children {
				stack.PushBack(v)
			}
			visited = append(visited, top)
		}
	}

	return nil
}
