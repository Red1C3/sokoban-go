package main

import (
	"sokoban/algor"
	"sokoban/game"
)

func main() {
	g:=game.NewGame("./puzzles/lab 1.json",&algor.NoAlgor{})
	g.Play()
}
