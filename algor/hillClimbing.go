package algor

import (
	"container/heap"
	"sokoban-go/state"
)

// HillClimbing implements itself.
type HillClimbing struct {
	steps int
}

func (hc *HillClimbing) Steps() int {
	return hc.steps
}

func (hc *HillClimbing) Search(start state.State) *state.State {
	visited := make([]state.State, 0)

	pq := make(state.PQ, 0)
	heap.Init(&pq)
	heap.Push(&pq, start)

	for len(pq) > 0 {
		hc.steps++
		min := heap.Pop(&pq).(state.State)
		if min.IsSolved() {
			return &min
		}

		children := min.States()
		for _, child := range children {
			if !isVisited(visited, child) {
				heap.Push(&pq, child)
				visited = append(visited, child)
			}
		}
		if len(pq) > 0 {
			min = heap.Pop(&pq).(state.State)
			pq = make(state.PQ, 0)
			heap.Push(&pq, min)
		} else {
			return nil
		}
	}

	return nil
}
