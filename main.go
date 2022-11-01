package main

import (
	"sokoban/state"
)

func main() {
	s := state.NewState("./puzzles/test.json")
	print(s.String())
}
