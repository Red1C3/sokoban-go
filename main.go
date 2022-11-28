/*
sokoban-go is an automated sokoban puzzles solver.

It provides easy to extend structures for implementing better searching algorithms
without having to design the whole program.

Usage:

	sokoban-go [algorithm] [puzzlepath]

algorithm options are:

	human
	dfs
	bfs
	astar-gh2
	astar-gh1
	astar-h1
	astar-h2
	hc-gh2
	hc-gh1
	hc-h2
	hc-h1

puzzlepath should be a suitable 2D JSON string array,
see "puzzles" folder for examples and [state/consts] for possible tiles characters.
*/
package main

import (
	"log"
	"os"
	"sokoban-go/algor"
	"sokoban-go/game"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("No enough command line arguments provided, please rtfm.")
	}
	switch os.Args[1] {
	case "dfs":
		g := game.NewGame(os.Args[2], &algor.Dfs{}, algor.NoFunc)
		g.Play()
	case "bfs":
		g := game.NewGame(os.Args[2], &algor.Bfs{}, algor.NoFunc)
		g.Play()
	case "astar-gh2":
		g := game.NewGame(os.Args[2], &algor.AStar{}, algor.GH2)
		g.Play()
	case "astar-gh1":
		g := game.NewGame(os.Args[2], &algor.AStar{}, algor.GH1)
		g.Play()
	case "astar-h2":
		g := game.NewGame(os.Args[2], &algor.AStar{}, algor.H2)
		g.Play()
	case "astar-h1":
		g := game.NewGame(os.Args[2], &algor.AStar{}, algor.H1)
		g.Play()
	case "hc-gh2":
		g := game.NewGame(os.Args[2], &algor.HillClimbing{}, algor.GH2)
		g.Play()
	case "hc-gh1":
		g := game.NewGame(os.Args[2], &algor.HillClimbing{}, algor.GH1)
		g.Play()
	case "hc-h2":
		g := game.NewGame(os.Args[2], &algor.HillClimbing{}, algor.H2)
		g.Play()
	case "hc-h1":
		g := game.NewGame(os.Args[2], &algor.HillClimbing{}, algor.H1)
		g.Play()
	case "human":
		g := game.NewGame(os.Args[2], &algor.NoAlgor{}, algor.NoFunc)
		g.Play()
	default:
		log.Fatalf("Unknown algorithm %s", os.Args[1])
	}
}
