package algor

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sokoban/state"
)

// NoAlgor is used for actual human input.
// it implements algor algorithm so it can be used
// the same as other algorithms.
type NoAlgor struct {
	steps    int
	curState state.State
}

func (a *NoAlgor) Steps() int {
	return a.steps
}

func (a *NoAlgor) Search(start state.State) *state.State {
	a.steps = 0
	a.curState = start
	println(a.curState.String())
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			log.Fatal("Failed to read input")
		}
		input := scanner.Text()
		nextStates := a.curState.StatesMap()
		var dir int
		switch input {
		case UP:
			dir = state.UP
		case DOWN:
			dir = state.DOWN
		case LEFT:
			dir = state.LEFT
		case RIGHT:
			dir = state.RIGHT
		default:
			continue
		}
		if s, ok := nextStates[dir]; ok {
			a.curState = s
			a.steps += 1
		}
		if a.curState.IsSolved() {
			break
		}
		println(a.curState.String())
	}
	return &a.curState
}
