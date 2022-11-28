package algor

import (
	"container/list"
	"sokoban-go/state"
)

// Bfs implements BFS algorithm.
type Bfs struct {
	steps int
}

func (b *Bfs) Steps() int {
	return b.steps
}

func (b *Bfs) Search(start state.State) *state.State {
	visited := make([]state.State, 0)

	queue := list.New()
	queue.PushBack(start)

	for queue.Len() > 0 {
		b.steps++
		first := queue.Remove(queue.Front()).(state.State)
		if first.IsSolved() {
			return &first
		}

		if !isVisited(visited, first) {
			children := first.States()
			for _, v := range children {
				queue.PushBack(v)
			}
			visited = append(visited, first)
		}
	}

	return nil
}
