package algor

import (
	"container/list"
	"sokoban/state"
)

type Bfs struct{

}

func (b *Bfs) Steps()int{
	return 0
}


func (b *Bfs)Search(start state.State)*state.State{
	visited := make([]state.State, 0)

	queue := list.New()
	queue.PushBack(start)

	for queue.Len() > 0 {
		first := queue.Remove(queue.Front()).(state.State)
		if first.IsSolved() {
			return &first
		}

		if !isVisited(visited, first) {
			children := first.States()
			for _, v := range children {
				queue.PushBack(v)
			}
			visited=append(visited, first)
		}
	}
	
	return nil
}