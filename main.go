package main

import (
	"sokoban/algor"
	"sokoban/game"
)

func main() {
	g:=game.NewGame("./puzzles/hard.json",&algor.HillClimbing{},algor.H2)
	g.Play()
}