package cell

// Cell represents a location in Game's grid.
type Cell struct {
	X, Y int
}

const (
	Up = iota
	Down
	Left
	Right
	None
)
