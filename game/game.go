package game

import (
	"sokoban/algor"
	"sokoban/state"
)

type game struct {
	algor      algor.Algor
	puzzlePath string
}

func NewGame(puzzle string, algor algor.Algor) game {
	var g game
	g.puzzlePath = puzzle
	g.algor = algor
	return g
}

func (g *game) Play() {
	start := state.NewState(g.puzzlePath)
	final := g.algor.Search(start)
	if final != nil {
		path := final.Path()
		for _, s := range path {
			println(s.String())
		}
		print("YAY")
	} else {
		print("No solution was found")
	}

}
