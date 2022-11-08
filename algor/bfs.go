package algor

import (
	"sokoban/state"
)

type Bfs struct{

}

func (b *Bfs) Steps()int{
	return 0
}


func (b *Bfs)Search(start state.State)state.State{
	visited := make([]state.State, 0)
	queue := make([]state.State, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		if first.IsSolved() {
			return first
		}

		if !isVisited(visited, first) {
			children := first.States()
			for _, v := range children {
				queue =append(queue,v)
			}
			visited=append(visited, first)
		}
	}
	return state.State{}
}