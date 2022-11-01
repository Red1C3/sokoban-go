package state

type state struct{
	tiles [][]int
	playerPos [2]int
}

func NewState()state{
	return state{}
}