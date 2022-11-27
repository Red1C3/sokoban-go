package game

import (
	"fmt"
	"sokoban/algor"
	"sokoban/state"
)

type game struct {
	algor         algor.Algor
	puzzlePath    string
	heuristicFunc func(*state.State) *int
}

func NewGame(puzzle string, algor algor.Algor, heuristicFunc func(*state.State) *int) game {
	var g game
	g.puzzlePath = puzzle
	g.algor = algor
	g.heuristicFunc = heuristicFunc
	return g
}

func (g *game) Play() {
	start := state.NewState(g.puzzlePath, g.heuristicFunc)
	final := g.algor.Search(start)
	if final != nil {
		if _, ok := g.algor.(*algor.NoAlgor); ok {
			println(final.String())
		} else {
			path := final.Path()
			for _, s := range path {
				println(s.String())
			}
		}
		println("YAY")
		fmt.Printf("Moves = %d", final.Moves)
	} else {
		print("No solution was found")
	}

}
