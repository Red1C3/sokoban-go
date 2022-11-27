// Package state implements Sokoban state structure and provides other helper functions.
package state

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

// State is the main structure that stores the
// Sokoban board in Tiles,
// and the number of moves to get to this state
// in Moves.
type State struct {
	Tiles     [][]int
	playerPos [2]int
	cost      *int
	costFunc  func(*State) *int
	parent    *State
	Moves     int
}

// Cost returns the cost of a state
func (s *State) Cost() *int {
	return s.cost
}

// NewState takes a the puzzle file path in puzzlePath,
// and the cost function in costFunc, and returns a state
// (mostly a start state).
func NewState(puzzlePath string, costFunc func(*State) *int) State {
	var s State
	puzzleFile, err := os.Open(puzzlePath)
	if err != nil {
		log.Fatalf("Failed to open file %s", puzzlePath)
	}
	decoder := json.NewDecoder(puzzleFile)
	var puzzleString [][]string
	err = decoder.Decode(&puzzleString)
	if err != nil {
		log.Fatalf("Failed to decode json file %s", puzzlePath)
	}
	s.Tiles = make([][]int, len(puzzleString)+BORDER*2)
	maxLength := len(puzzleString[0])

	for _, a := range puzzleString {
		if len(a) > maxLength {
			maxLength = len(a)
		}
	}

	for i := 0; i < len(s.Tiles); i++ {
		s.Tiles[i] = make([]int, maxLength+2*BORDER)
		for j := 0; j < len(s.Tiles[i]); j++ {
			s.Tiles[i][j] = OBSTACLE
		}
	}

	for i, a := range puzzleString {
		for j, v := range a {
			switch v {
			case BLANKCHAR:
				s.Tiles[i+BORDER][j+BORDER] = BLANK
			case BOXCHAR:
				s.Tiles[i+BORDER][j+BORDER] = BOX
			case OBSTACLECHAR:
				s.Tiles[i+BORDER][j+BORDER] = OBSTACLE
			case PLAYERCHAR:
				s.Tiles[i+BORDER][j+BORDER] = PLAYER
				s.playerPos = [2]int{i + BORDER, j + BORDER}
			case GOALCHAR:
				s.Tiles[i+BORDER][j+BORDER] = GOAL
			case BOXONGOALCHAR:
				s.Tiles[i+BORDER][j+BORDER] = BOXONGOAL
			case PLAYERONGOALCHAR:
				s.Tiles[i+BORDER][j+BORDER] = PLAYERONGOAL
				s.playerPos = [2]int{i + BORDER, j + BORDER}
			default:
				log.Fatalf("Unknown puzzle char %s", v)
			}
		}
	}
	s.costFunc = costFunc
	s.Moves = 0
	return s
}

// String returns a human-readable representation of s.
func (s *State) String() string {
	var buffer bytes.Buffer
	for _, a := range s.Tiles {
		for _, v := range a {
			switch v {
			case BLANK:
				buffer.WriteString(BLANKCHAR)
			case OBSTACLE:
				buffer.WriteString(OBSTACLECHAR)
			case BOX:
				buffer.WriteString(BOXCHAR)
			case PLAYER:
				buffer.WriteString(PLAYERCHAR)
			case GOAL:
				buffer.WriteString(GOALCHAR)
			case BOXONGOAL:
				buffer.WriteString(BOXONGOALCHAR)
			case PLAYERONGOAL:
				buffer.WriteString(PLAYERONGOALCHAR)
			default:
				log.Fatalf("Unknown puzzle digit %d", v)
			}
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

// Equlas checks if s is identical to o.
func (s *State) Equals(o *State) bool {
	for i, a := range s.Tiles {
		for j := range a {
			if s.Tiles[i][j] != o.Tiles[i][j] { //Does not compare parent
				return false
			}
		}
	}
	return true
}

func (s *State) canMove(dir int) bool {
	next := [2]int{0, 0}
	switch dir {
	case UP:
		next[0] = s.Tiles[s.playerPos[0]-1][s.playerPos[1]]
		if (next[0] & BOX) != 0 {
			next[1] = s.Tiles[s.playerPos[0]-2][s.playerPos[1]]
		}
	case DOWN:
		next[0] = s.Tiles[s.playerPos[0]+1][s.playerPos[1]]
		if (next[0] & BOX) != 0 {
			next[1] = s.Tiles[s.playerPos[0]+2][s.playerPos[1]]
		}
	case LEFT:
		next[0] = s.Tiles[s.playerPos[0]][s.playerPos[1]-1]
		if (next[0] & BOX) != 0 {
			next[1] = s.Tiles[s.playerPos[0]][s.playerPos[1]-2]
		}
	case RIGHT:
		next[0] = s.Tiles[s.playerPos[0]][s.playerPos[1]+1]
		if (next[0] & BOX) != 0 {
			next[1] = s.Tiles[s.playerPos[0]][s.playerPos[1]+2]
		}
	}
	if (next[0] & BOX) == 0 {
		return (next[0]^BLANK == 0) || (next[0]^GOAL == 0)
	} else {
		return (next[1]^BLANK == 0) || (next[1]^GOAL == 0)
	}
}

func (s *State) makeCopy() State {
	newState := State{playerPos: s.playerPos,
		cost:     nil,
		costFunc: s.costFunc}

	newState.Tiles = make([][]int, len(s.Tiles))
	for i, arr := range s.Tiles {
		newState.Tiles[i] = make([]int, len(arr))
		copy(newState.Tiles[i], arr)
	}
	return newState
}

// move assumes canMove is true
func (s *State) move(dir int) State {
	newState := s.makeCopy()
	switch dir {
	case UP:
		if (newState.Tiles[newState.playerPos[0]-1][newState.playerPos[1]] & BOX) != 0 {
			newState.Tiles[newState.playerPos[0]-2][newState.playerPos[1]] |= BOX
			newState.Tiles[newState.playerPos[0]-1][newState.playerPos[1]] &= ^BOX
		}
		newState.Tiles[newState.playerPos[0]-1][newState.playerPos[1]] |= PLAYER
		newState.Tiles[newState.playerPos[0]][newState.playerPos[1]] &= ^PLAYER
		newState.playerPos[0] -= 1
	case DOWN:
		if (newState.Tiles[newState.playerPos[0]+1][newState.playerPos[1]] & BOX) != 0 {
			newState.Tiles[newState.playerPos[0]+2][newState.playerPos[1]] |= BOX
			newState.Tiles[newState.playerPos[0]+1][newState.playerPos[1]] &= ^BOX
		}
		newState.Tiles[newState.playerPos[0]+1][newState.playerPos[1]] |= PLAYER
		newState.Tiles[newState.playerPos[0]][newState.playerPos[1]] &= ^PLAYER
		newState.playerPos[0] += 1
	case LEFT:
		if (newState.Tiles[newState.playerPos[0]][newState.playerPos[1]-1] & BOX) != 0 {
			newState.Tiles[newState.playerPos[0]][newState.playerPos[1]-2] |= BOX
			newState.Tiles[newState.playerPos[0]][newState.playerPos[1]-1] &= ^BOX
		}
		newState.Tiles[newState.playerPos[0]][newState.playerPos[1]-1] |= PLAYER
		newState.Tiles[newState.playerPos[0]][newState.playerPos[1]] &= ^PLAYER
		newState.playerPos[1] -= 1
	case RIGHT:
		if (newState.Tiles[newState.playerPos[0]][newState.playerPos[1]+1] & BOX) != 0 {
			newState.Tiles[newState.playerPos[0]][newState.playerPos[1]+2] |= BOX
			newState.Tiles[newState.playerPos[0]][newState.playerPos[1]+1] &= ^BOX
		}
		newState.Tiles[newState.playerPos[0]][newState.playerPos[1]+1] |= PLAYER
		newState.Tiles[newState.playerPos[0]][newState.playerPos[1]] &= ^PLAYER
		newState.playerPos[1] += 1
	}
	newState.parent = s
	newState.costFunc = s.costFunc
	newState.Moves = s.Moves + 1
	return newState
}

// States returns all the possible states from s
func (s *State) States() []State {
	states := make([]State, 0)
	if s.canMove(UP) {
		states = append(states, s.move(UP))
	}
	if s.canMove(DOWN) {
		states = append(states, s.move(DOWN))
	}
	if s.canMove(RIGHT) {
		states = append(states, s.move(RIGHT))
	}
	if s.canMove(LEFT) {
		states = append(states, s.move(LEFT))
	}
	return states
}

// StatesMap maps the possible states from s to
// their direction enum (UP/DOWN/LEFT/RIGHT)
func (s *State) StatesMap() map[int]State {
	statesMap := make(map[int]State)
	if s.canMove(UP) {
		statesMap[UP] = s.move(UP)
	}
	if s.canMove(DOWN) {
		statesMap[DOWN] = s.move(DOWN)
	}
	if s.canMove(LEFT) {
		statesMap[LEFT] = s.move(LEFT)
	}
	if s.canMove(RIGHT) {
		statesMap[RIGHT] = s.move(RIGHT)
	}
	return statesMap
}

// IsSolved checks if s is a solution state
func (s *State) IsSolved() bool {
	for _, a := range s.Tiles {
		for _, v := range a {
			if (v^GOAL == 0) || (v^PLAYERONGOAL == 0) {
				return false
			}
		}
	}
	return true
}

// Path returns the path from the oldest ancestor of s
// to s itself
func (s *State) Path() []State {
	path := make([]State, 0)
	cur := s
	for cur != nil {
		path = append(path, *cur)
		cur = cur.parent
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

// Pos returns the player position (x,y) in s
func (s *State) Pos() [2]int {
	return s.playerPos
}
