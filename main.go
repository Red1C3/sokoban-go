package main

import (
	"sokoban/state"
)

func main() {
	s := state.NewState("./puzzles/test 1.json")
	print(s.States()[0].String())
}
