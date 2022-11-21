package algor

import (
	"container/heap"
	"sokoban/state"
)

type AStar struct{
	steps int
}

func (a *AStar) Steps()int {
	return a.steps
}

func (a *AStar)Search(start state.State)*state.State{
	visited := make([]state.State, 0)

	pq:=make(state.PQ,0)
	heap.Init(&pq)
	heap.Push(&pq,start)

	for len(pq) > 0 {
		a.steps++
		min:=heap.Pop(&pq).(state.State)
		if min.IsSolved() {
			return &min
		}
		if !isVisited(visited, min) {
			children := min.States()
			for _, v := range children {
				heap.Push(&pq,v)
			}
			visited=append(visited, min)
		}
	}

	return nil
}