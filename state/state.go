package state

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strings"
)

type state struct {
	tiles     [][]int
	playerPos [2]int
}

func NewState(puzzlePath string) state {
	var s state
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
			default:
				log.Fatalf("Unknown puzzle char %s", v)
			}
		}
	}
	return s
}

func (s *state) String() string {
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
			default:
				log.Fatalf("Unknown puzzle digit %d", v)
			}
		}
		buffer.WriteString("\n" + strings.Repeat(HSEPERATOR, len(a)*2+1) + "\n")
	}
	return buffer.String()
}

func (s *state) Equals(o *state) bool {
	for i, a := range s.tiles {
		for j := range a {
			if s.tiles[i][j] != o.tiles[i][j] {
				return false
			}
		}
	}
	return true
}

func (s *state) canMove(dir int) bool {
	next := [2]int{0, 0}
	switch dir {
	case UP:
		next[0] = s.tiles[s.playerPos[0]-1][s.playerPos[1]]
		if next[0] == BOX {
			next[1] = s.tiles[s.playerPos[0]-2][s.playerPos[1]]
		}
	case DOWN:
		next[0] = s.tiles[s.playerPos[0]+1][s.playerPos[1]]
		if next[0] == BOX {
			next[1] = s.tiles[s.playerPos[0]+2][s.playerPos[1]]
		}
	case LEFT:
		next[0] = s.tiles[s.playerPos[0]][s.playerPos[1]-1]
		if next[0] == BOX {
			next[1] = s.tiles[s.playerPos[0]][s.playerPos[1]-2]
		}
	case RIGHT:
		next[0] = s.tiles[s.playerPos[0]][s.playerPos[1]+1]
		if next[0] == BOX {
			next[1] = s.tiles[s.playerPos[0]][s.playerPos[1]+2]
		}
	}
	if next[0] != BOX {
		return next[0] == BLANK
	} else {
		return (next[0] ^ next[1]) == (BOX ^ BLANK)
	}
}

func (s *state) makeCopy() state {
	newState := state{playerPos: s.playerPos}
	newState.tiles = make([][]int, len(s.tiles))
	for i, arr := range s.tiles {
		newState.tiles[i] = make([]int, len(arr))
		copy(newState.tiles[i], arr)
	}
	return newState
}

func (s *state) move(dir int) (state, bool) {
	if s.canMove(dir) {
		newState := s.makeCopy()
		switch dir {
		case UP:
			if newState.tiles[newState.playerPos[0]-1][newState.playerPos[1]] == BOX {
				newState.tiles[newState.playerPos[0]-2][newState.playerPos[1]] = BOX
			}
			newState.tiles[newState.playerPos[0]-1][newState.playerPos[1]] = PLAYER
			newState.tiles[newState.playerPos[0]][newState.playerPos[1]] = BLANK
			newState.playerPos[0] -= 1
		case DOWN:
			if newState.tiles[newState.playerPos[0]+1][newState.playerPos[1]] == BOX {
				newState.tiles[newState.playerPos[0]+2][newState.playerPos[1]] = BOX
			}
			newState.tiles[newState.playerPos[0]+1][newState.playerPos[1]] = PLAYER
			newState.tiles[newState.playerPos[0]][newState.playerPos[1]] = BLANK
			newState.playerPos[0] += 1
		case LEFT:
			if newState.tiles[newState.playerPos[0]][newState.playerPos[1]-1] == BOX {
				newState.tiles[newState.playerPos[0]][newState.playerPos[1]-2] = BOX
			}
			newState.tiles[newState.playerPos[0]][newState.playerPos[1]-1] = PLAYER
			newState.tiles[newState.playerPos[0]][newState.playerPos[1]] = BLANK
			newState.playerPos[1] -= 1
		case RIGHT:
			if newState.tiles[newState.playerPos[0]][newState.playerPos[1]+1] == BOX {
				newState.tiles[newState.playerPos[0]][newState.playerPos[1]+2] = BOX
			}
			newState.tiles[newState.playerPos[0]][newState.playerPos[1]+1] = PLAYER
			newState.tiles[newState.playerPos[0]][newState.playerPos[1]] = BLANK
			newState.playerPos[1] += 1
		}
		return newState, true
	} else {
		return state{}, false
	}
}
