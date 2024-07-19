package shared

type GameState struct {
	Grid      [8][7]int8
	playerOne int8
	playerTwo int8
}

type Logic interface {
	Init(*GameState)
	Close()
	GetAvailableMoves() ([4]int8, error)
	PlayMove(int8) error
}

type Renderer interface {
	Init(*GameState)
	Close()
	Render() error
}

const (
	red    = iota // 0
	purple = iota // 1
	blue   = iota // 2
	green  = iota // 3
	yellow = iota // 4
	black  = iota // 5
)
