package main

import (
	"sokoban/algor"
	"sokoban/state"
)

func main() {
	s := state.NewState("./puzzles/lab 1.json")
	algor:=algor.NoAlgor{}
	final:=algor.Search(s)
	print(final.String())
}
