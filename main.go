package main

import (
	"sokoban/algor"
	"sokoban/game"
)

func main() {
	g:=game.NewGame("./puzzles/simple.json",&algor.AStar{},algor.H1)
	g.Play()
}