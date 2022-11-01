package main

import (
	"sokoban/algor"
	"sokoban/game"
)

func main() {
	g:=game.NewGame("./puzzles/test 1.json",&algor.NoAlgor{})
	g.Play()
}
