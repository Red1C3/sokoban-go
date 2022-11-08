package main

import (
	"sokoban/algor"
	"sokoban/game"
)

func main() {
	g:=game.NewGame("./puzzles/simple.json",&algor.Bfs{},algor.NoFunc)
	g.Play()
}