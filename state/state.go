package state

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strings"
)

type State struct {
	tiles     [][]int
	playerPos [2]int
}

func NewState(puzzlePath string) State {
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
	s.tiles = make([][]int, len(puzzleString)+BORDER*2)
	maxLength := len(puzzleString[0])

	for _, a := range puzzleString {
		if len(a) > maxLength {
			maxLength = len(a)
		}
	}

	for i := 0; i < len(s.tiles); i++ {
		s.tiles[i] = make([]int, maxLength+2*BORDER)
		for j := 0; j < len(s.tiles[i]); j++ {
			s.tiles[i][j] = OBSTACLE
		}
	}

	for i, a := range puzzleString {
		for j, v := range a {
			switch v {
			case BLANKCHAR:
				s.tiles[i+BORDER][j+BORDER] = BLANK
			case BOXCHAR:
				s.tiles[i+BORDER][j+BORDER] = BOX
			case OBSTACLECHAR:
				s.tiles[i+BORDER][j+BORDER] = OBSTACLE
			case PLAYERCHAR:
				s.tiles[i+BORDER][j+BORDER] = PLAYER
				s.playerPos = [2]int{i + 1, j + 1}
			case GOALCHAR:
				s.tiles[i+BORDER][j+BORDER] = GOAL
			case BOXONGOALCHAR:
				s.tiles[i+BORDER][j+BORDER] = BOXONGOAL
			case PLAYERONGOALCHAR:
				s.tiles[i+BORDER][j+BORDER] = PLAYERONGOAL
			default:
				log.Fatalf("Unknown puzzle char %s", v)
			}
		}
	}
	return s
}

func (s *State) String() string {
	var buffer bytes.Buffer
	for _, a := range s.tiles {
		buffer.WriteString(VSEPERATOR)
		for _, v := range a {
			switch v {
			case BLANK:
				buffer.WriteString(BLANKCHAR + VSEPERATOR)
			case OBSTACLE:
				buffer.WriteString(OBSTACLECHAR + VSEPERATOR)
			case BOX:
				buffer.WriteString(BOXCHAR + VSEPERATOR)
			case PLAYER:
				buffer.WriteString(PLAYERCHAR + VSEPERATOR)
			case GOAL:
				buffer.WriteString(GOALCHAR + VSEPERATOR)
			case BOXONGOAL:
				buffer.WriteString(BOXONGOALCHAR + VSEPERATOR)
			case PLAYERONGOAL:
				buffer.WriteString(PLAYERONGOALCHAR + VSEPERATOR)
			default:
				log.Fatalf("Unknown puzzle digit %d", v)
			}
		}
		buffer.WriteString("\n" + strings.Repeat(HSEPERATOR, len(a)*3+1) + "\n")
	}
	return buffer.String()
}

func (s *State) Equals(o *State) bool {
	for i, a := range s.tiles {
		for j := range a {
			if s.tiles[i][j] != o.tiles[i][j] {
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
		next[0] = s.tiles[s.playerPos[0]-1][s.playerPos[1]]
		if (next[0] & BOX) != 0 {
			next[1] = s.tiles[s.playerPos[0]-2][s.playerPos[1]]
		}
	case DOWN:
		next[0] = s.tiles[s.playerPos[0]+1][s.playerPos[1]]
		if (next[0] & BOX) != 0 {
			next[1] = s.tiles[s.playerPos[0]+2][s.playerPos[1]]
		}
	case LEFT:
		next[0] = s.tiles[s.playerPos[0]][s.playerPos[1]-1]
		if (next[0] & BOX) != 0 {
			next[1] = s.tiles[s.playerPos[0]][s.playerPos[1]-2]
		}
	case RIGHT:
		next[0] = s.tiles[s.playerPos[0]][s.playerPos[1]+1]
		if (next[0] & BOX) != 0 {
			next[1] = s.tiles[s.playerPos[0]][s.playerPos[1]+2]
		}
	}
	if (next[0] & BOX) == 0 {
		return (next[0]^BLANK == 0) || (next[0]^GOAL == 0)
	} else {
		return (next[1]^BLANK == 0) || (next[1]^GOAL == 0)
	}
}

func (s *State) makeCopy() State {
	newState := State{playerPos: s.playerPos}
	newState.tiles = make([][]int, len(s.tiles))
	for i, arr := range s.tiles {
		newState.tiles[i] = make([]int, len(arr))
		copy(newState.tiles[i], arr)
	}
	return newState
}

// Assumes canMove is true
func (s *State) move(dir int) State {
	newState := s.makeCopy()
	switch dir {
	case UP:
		if (newState.tiles[newState.playerPos[0]-1][newState.playerPos[1]] & BOX) != 0 {
			newState.tiles[newState.playerPos[0]-2][newState.playerPos[1]] |= BOX
			newState.tiles[newState.playerPos[0]-1][newState.playerPos[1]] &= ^BOX
		}
		newState.tiles[newState.playerPos[0]-1][newState.playerPos[1]] |= PLAYER
		newState.tiles[newState.playerPos[0]][newState.playerPos[1]] &= ^PLAYER
		newState.playerPos[0] -= 1
	case DOWN:
		if (newState.tiles[newState.playerPos[0]+1][newState.playerPos[1]] & BOX) != 0 {
			newState.tiles[newState.playerPos[0]+2][newState.playerPos[1]] |= BOX
			newState.tiles[newState.playerPos[0]+1][newState.playerPos[1]] &= ^BOX
		}
		newState.tiles[newState.playerPos[0]+1][newState.playerPos[1]] |= PLAYER
		newState.tiles[newState.playerPos[0]][newState.playerPos[1]] &= ^PLAYER
		newState.playerPos[0] += 1
	case LEFT:
		if (newState.tiles[newState.playerPos[0]][newState.playerPos[1]-1] & BOX) != 0 {
			newState.tiles[newState.playerPos[0]][newState.playerPos[1]-2] |= BOX
			newState.tiles[newState.playerPos[0]][newState.playerPos[1]-1] &= ^BOX
		}
		newState.tiles[newState.playerPos[0]][newState.playerPos[1]-1] |= PLAYER
		newState.tiles[newState.playerPos[0]][newState.playerPos[1]] &= ^PLAYER
		newState.playerPos[1] -= 1
	case RIGHT:
		if (newState.tiles[newState.playerPos[0]][newState.playerPos[1]+1] & BOX) != 0 {
			newState.tiles[newState.playerPos[0]][newState.playerPos[1]+2] |= BOX
			newState.tiles[newState.playerPos[0]][newState.playerPos[1]+1] &= ^BOX
		}
		newState.tiles[newState.playerPos[0]][newState.playerPos[1]+1] |= PLAYER
		newState.tiles[newState.playerPos[0]][newState.playerPos[1]] &= ^PLAYER
		newState.playerPos[1] += 1
	}
	return newState
}

func (s *State) States() []State {
	states := make([]State, 0)
	if s.canMove(UP) {
		states = append(states, s.move(UP))
	}
	if s.canMove(DOWN) {
		states = append(states, s.move(DOWN))
	}
	if s.canMove(LEFT) {
		states = append(states, s.move(LEFT))
	}
	if s.canMove(RIGHT) {
		states = append(states, s.move(RIGHT))
	}
	return states
}

func (s *State) IsSolved() bool {
	for _, a := range s.tiles {
		for _, v := range a {
			if v^GOAL == 0 {
				return false
			}
		}
	}
	return true
}
