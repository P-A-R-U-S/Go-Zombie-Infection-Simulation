package Data

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
	maze          [][]bool
	height, width int
}

func NewWorld() *World {
	maze1 := [][]bool{
		//0     1      2      3      4      5      6      7      8      9      10     11     12     13     14     15     16     17    18     19     20     21     22     23     24     25     26     27     28
		{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},                                     // 0
		{true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, true, false, false, false},   // 1
		{true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, true, true, true, true, true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, true, true, true, true, true, false, true, false, true, true, true, true, true},                         // 2
		{true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, true, true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, true, true, false, true, true, true, false, true, true, true},                         // 3
		{true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true},           // 4
		{false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true},                 // 5
		{true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true},                 // 6
		{true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, false, true, true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, false, true, true, true, false, false, true, false, true, true, true},                 // 7
		{true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, false, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, false, true, true, true, false, true, true, false, true, true, true},                       // 8
		{true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, true},                                 // 9
		{true, true, true, false, false, true, false, false, false, false, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, false, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, true, false, false, false}, // 10
		{true, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, true, false, true, false, true, true, true, true, true},                       // 11
		{true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, true, true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, true, true, false, true, true, true, false, true, true, true},                         // 12
		{true, false, true, true, false, false, true, false, true, true, false, true, false, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, true, false, true, false, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true},         // 13
		{false, false, true, true, true, true, true, false, true, false, false, true, false, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, false, true, false, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true},               // 14
		{true, true, true, true, true, false, false, false, true, true, false, true, false, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, true, false, true, false, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true},             // 15
		{true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, false, false, true, false, true, true, true},                   // 16
		{true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true},                         // 17
		{true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true},                         // 18
		{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true},                                   // 19
		{true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, false, false, false, false, true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, false, false, false, false}, // 20
		{true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, false, true, true, true, true, false, true, false, true, false, true, true, true, true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, false, true, true, true, true, false, true, false, true, false, true, true, true},                     // 21
		{true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, false, true, true, true, true, false, true, true, true, false, true, true, true},                       // 22
		{true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true},           // 23
		{false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true},                 // 24
		{true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true},                 // 25
		{true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, false, false, true, false, true, true, true},                   // 26
		{true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true},                         // 27
		{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},                                     // 0
		{true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, true, false, false, false},   // 1
		{true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, true, true, true, true, true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, true, true, true, true, true, false, true, false, true, true, true, true, true},                         // 2
		{true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, true, true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, true, true, false, true, true, true, false, true, true, true},                         // 3
		{true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true},           // 4
		{false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true},                 // 5
		{true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true},                 // 6
		{true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, false, true, true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, false, true, true, true, false, false, true, false, true, true, true},                 // 7
		{true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, false, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, false, true, true, true, false, true, true, false, true, true, true},                       // 8
		{true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, true},                                 // 9
		{true, true, true, false, false, true, false, false, false, false, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, false, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, true, false, false, false}, // 10
		{true, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, true, false, true, false, true, true, true, true, true},                       // 11
		{true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, true, true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, true, true, false, true, true, true, false, true, true, true},                         // 12
		{true, false, true, true, false, false, true, false, true, true, false, true, false, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, true, false, true, false, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true},         // 13
		{false, false, true, true, true, true, true, false, true, false, false, true, false, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, false, true, false, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true},               // 14
		{true, true, true, true, true, false, false, false, true, true, false, true, false, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, true, false, true, false, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true},             // 15
		{true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, false, false, true, false, true, true, true},                   // 16
		{true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true},                         // 17
		{true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true},                         // 18
		{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true},                                   // 19
		{true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, false, false, false, false, true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, false, false, false, false}, // 20
		{true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, false, true, true, true, true, false, true, false, true, false, true, true, true, true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, false, true, true, true, true, false, true, false, true, false, true, true, true},                     // 21
		{true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, false, true, true, true, true, false, true, true, true, false, true, true, true},                       // 22
		{true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true},           // 23
		{false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true},                 // 24
		{true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true},                 // 25
		{true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, false, false, true, false, true, true, true},                   // 26
		{true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true},                         // 27

	}

	return &World{maze: maze1,
		height: len(maze1) - 1,
		width:  len(maze1[0]) - 1}
}
