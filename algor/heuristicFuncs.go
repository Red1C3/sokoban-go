package algor

import (
	"math"
	"sokoban/state"
)

var NoFunc = func(s *state.State) *int {
	return nil
}

var UniformCost=func(s *state.State)*int{
	return &s.Moves
}

var H1 = func(s *state.State) *int {
	unsolvedBoxes := make([][2]int, 0)
	unsolvedGoals := make([][2]int, 0)
	for i, a := range s.Tiles {
		for j, v := range a {
			if v == state.BOX {
				unsolvedBoxes = append(unsolvedBoxes, [2]int{i, j})
			}
			if v == state.GOAL {
				unsolvedGoals = append(unsolvedGoals, [2]int{i, j})
			}
		}
	}
	heuristic := 0
	for _, g := range unsolvedGoals {
		minDist := math.MaxInt
		for _, b := range unsolvedBoxes {
			minDist = int(math.Min(float64(minDist), float64(manhattanDist(g, b))))
		}
		heuristic += minDist
	}
	heuristicNCost:=heuristic+s.Moves
	return &heuristicNCost
}

var H2 = func(s *state.State) *int {
	unsolvedBoxes := make([][2]int, 0)
	unsolvedGoals := make([][2]int, 0)
	for i, a := range s.Tiles {
		for j, v := range a {
			if v == state.BOX {
				unsolvedBoxes = append(unsolvedBoxes, [2]int{i, j})
			}
			if v == state.GOAL {
				unsolvedGoals = append(unsolvedGoals, [2]int{i, j})
			}
		}
	}
	heuristic := 0
	for _, g := range unsolvedGoals {
		minDist := math.MaxInt
		for _, b := range unsolvedBoxes {
			minDist = int(math.Min(float64(minDist), float64(manhattanDist(g, b))))
		}
		heuristic += minDist
	}

	player:=s.Pos()

	minDist:=math.MaxInt

	for _,b:=range unsolvedBoxes{
		minDist=int(math.Min(float64(minDist),float64(manhattanDist(player,b))))
	}

	heuristic+=minDist
	heuristicNCost:=heuristic+s.Moves
	return &heuristicNCost
}

func manhattanDist(a, b [2]int) int {
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}
