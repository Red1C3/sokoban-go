package main

import (
	"sokoban/algor"
	"sokoban/game"
)

func main() {
	g:=game.NewGame("./puzzles/hard.json",&algor.AStar{},algor.H2)
	g.Play()
}