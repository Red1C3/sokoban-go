// Package game is just a wrapper for the searching process.
package game

import (
	"fmt"
	"sokoban-go/algor"
	"sokoban-go/state"
)

type game struct {
	algor      algor.Algor
	puzzlePath string
	costFunc   func(*state.State) *int
}

// NewGame creates a new game instance.
//
// it takes the start state path in puzzle, the desired algorithm in
// algor, and the desired cost function in costFunc.
//
// In case the algorithm doesn't require a cost function, algor.NoFunc
// can be passed, but any other valid functions will work too
// and will be simply ignored.
func NewGame(puzzle string, algor algor.Algor, costFunc func(*state.State) *int) game {
	var g game
	g.puzzlePath = puzzle
	g.algor = algor
	g.costFunc = costFunc
	return g
}

// Play starts the searching algorithm process,
// or in case of algor.NoAlgor, the human player prompting.
//
// It prints the path to the final state and the length of that path
// minus the start state.
//
// In case no solution was found, it simply prints "No solution was found"
func (g *game) Play() {
	start := state.NewState(g.puzzlePath, g.costFunc)
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
