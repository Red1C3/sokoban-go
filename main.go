/*
sokoban-go is an automated sokoban puzzles solver.

It provides easy to extend structures for implementing better searching algorithms
without having to design the whole program.
*/
package main

import (
	"sokoban/algor"
	"sokoban/game"
)

func main() {
	g := game.NewGame("./puzzles/simple.json", &algor.Bfs{}, algor.GH2) //TODO take terminal args instead of hardcode
	g.Play()
}
