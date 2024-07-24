package Data

import (
	"fmt"
	"math/rand"
)

// Reset to default
const RESET = "\u001b[0m"

// Symbols Colors
const (
	RED     = "\u001b[31m"
	GREEN   = "\u001b[32m"
	BLACK   = "\u001b[30m"
	YELLOW  = "\u001b[33m"
	BLUE    = "\u001b[34m"
	MAGENTA = "\u001b[35m"
	CYAN    = "\u001b[36m"
	WHITE   = "\u001b[37m"

	BRIGHT_BLACK   = "\u001b[30;1m"
	BRIGHT_RED     = "\u001b[31;1m"
	BRIGHT_GREEN   = "\u001b[32;1m"
	BRIGHT_YELLOW  = "\u001b[33;1m"
	BRIGHT_BLUE    = "\u001b[34;1m"
	BRIGHT_MAGENTA = "\u001b[35;1m"
	BRIGHT_CYAN    = "\u001b[36;1m"
	BRIGHT_WHITE   = "\u001b[37;1m"
)

// Background Color
const (
	BACKGROUND_BLACK   = "\u001b[40m"
	BACKGROUND_RED     = "\u001b[41m"
	BACKGROUND_GREEN   = "\u001b[42m"
	BACKGROUND_YELLOW  = "\u001b[43m"
	BACKGROUND_BLUE    = "\u001b[44m"
	BACKGROUND_MAGENTA = "\u001b[45m"
	BACKGROUND_CYAN    = "\u001b[46m"
	BACKGROUND_WHITE   = "\u001b[47m"

	BACKGROUND_BRIGHT_BLACK   = "\u001b[40;1m"
	BACKGROUND_BRIGHT_RED     = "\u001b[41;1m"
	BACKGROUND_BRIGHT_GREEN   = "\u001b[42;1m"
	BACKGROUND_BRIGHT_YELLOW  = "\u001b[43;1m"
	BACKGROUND_BRIGHT_BLUE    = "\u001b[44;1m"
	BACKGROUND_BRIGHT_MAGENTA = "\u001b[45;1m"
	BACKGROUND_BRIGHT_CYAN    = "\u001b[46;1m"
	BACKGROUND_BRIGHT_WHITE   = "\u001b[47;1m"
)

// Symbols
const (
	START    = "*"
	END      = "X"
	VISITED  = "◼"
	PATH     = "◼"
	OBSTACLE = " "
	ENEMY    = ""
	SPACE    = " "
)

type World struct {
	maze          [][]*Cell
	cells         map[*Entity]*Cell
	height, width int
}

type Cell struct {
	entity *Entity
	prev   *Cell
	next   *Cell
}

func NewWorld() *World {
	maze1 := [][]*Cell{
		//0     1      2      3      4      5      6      7      8      9      10     11     12     13     14     15     16     17    18     19     20     21     22     23     24     25     26     27     28
		{&Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}}, // 0
		{&Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil},                                                                                                                                         // 1
		{&Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}},                                                 // 2
		{&Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                 // 3
		{&Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}},                                                                                                         // 4
		{nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}},                                                                                 // 5
		{&Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}},                                                                                 // 6
		{&Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                                                 // 7
		{&Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                         // 8
		{&Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}},                 // 9
		{&Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil},                                                                                                                                                 // 10
		{&Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}},                                                         // 11
		{&Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                 // 12
		{&Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}},                                                                                                                 // 13
		{nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}},                                                                                         // 14
		{&Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}},                                                                                                 // 15
		{&Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                                         // 16
		{&Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                 // 17
		{&Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                 // 18
		{&Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},         // 19
		{&Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, nil, nil, nil},                                                                                                                                                 // 20
		{&Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                                 // 21
		{&Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                         // 22
		{&Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}},                                                                                                         // 23
		{nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}},                                                                                 // 24
		{&Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}},                                                                                 // 25
		{&Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                                         // 26
		{&Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                 // 27
		{&Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}}, // 0
		{&Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil},                                                                                                                                         // 1
		{&Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}},                                                 // 2
		{&Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                 // 3
		{&Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}},                                                                                                         // 4
		{nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}},                                                                                 // 5
		{&Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}},                                                                                 // 6
		{&Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                                                 // 7
		{&Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                         // 8
		{&Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}},                 // 9
		{&Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil},                                                                                                                                                 // 10
		{&Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}},                                                         // 11
		{&Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                 // 12
		{&Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}},                                                                                                                 // 13
		{nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}},                                                                                         // 14
		{&Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}},                                                                                                 // 15
		{&Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                                         // 16
		{&Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                 // 17
		{&Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                 // 18
		{&Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},         // 19
		{&Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, nil, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, nil, nil, nil},                                                                                                                                                 // 20
		{&Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                                 // 21
		{&Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                         // 22
		{&Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}},                                                                                                         // 23
		{nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, nil, &Cell{}, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}},                                                                                 // 24
		{&Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, nil, &Cell{}},                                                                                 // 25
		{&Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, nil, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                                         // 26
		{&Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, nil, &Cell{}, &Cell{}, &Cell{}},                                                 // 27
	}

	return &World{maze: maze1,
		height: len(maze1) - 1,
		width:  len(maze1[0]) - 1,
		cells:  map[*Entity]*Cell{},
	}
}

func (w *World) SetToCell(entity *Entity) {
	cell := &Cell{}
	cell.entity = entity

	var isFound bool
	for !isFound {
		entity.x = rand.Intn(w.width - 1)
		entity.y = rand.Intn(w.height - 1)

		isFound = w.maze[entity.x][entity.y] != nil
	}

	tail := w.maze[cell.entity.x][cell.entity.y]
	for tail.next != nil {
		tail = tail.next
	}
	tail.next, cell.prev = cell, tail
}

func (w *World) Draw() {
	const RESET = "\u001b[0m"

	ROWS := len(w.maze)
	COLS := len(w.maze[0])

	// Clean
	//fmt.Print("\u001b[1000D")                         //Move left
	//fmt.Print("\u001b[" + strconv.Itoa(ROWS+8) + "A") // Move up

	for row := 0; row < ROWS; row++ {
		//fmt.Print(RESET, fmt.Sprintf("%4d|", row))
		for col := 0; col < COLS; col++ {
			fmt.Print(RESET)

			if w.maze[row][col] == nil {
				fmt.Print(BACKGROUND_WHITE, SPACE, OBSTACLE, SPACE)
				continue
			}

			if w.maze[row][col].next == nil {
				fmt.Print(SPACE, SPACE, SPACE)
				continue
			}

			fmt.Print(SPACE)

			isZombieInCell := false
			h := w.maze[row][col].next
			for h != nil && !isZombieInCell {
				isZombieInCell = h.entity.entityType == Zombie
				h = h.next
			}

			if isZombieInCell {
				fmt.Print(RED, VISITED)
			} else {
				fmt.Print(GREEN, PATH)
			}
			fmt.Print(SPACE)
			continue
		}
		fmt.Println(RESET)
	}
	//fmt.Println(RESET)
	//fmt.Println()

	//fmt.Println(Background_Yellow, " ", RESET, " - In Stack")
	//fmt.Println(RED, VISITED, RESET, " - Visited")
	//fmt.Println(Background_Bright_White, OBSTACLE, RESET, " - Obstacle")
	//fmt.Println(GREEN, VISITED, RESET, " - Path")

	//Sleep(10)
}
