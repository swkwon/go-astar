package astar

type Node struct {
	Coord   Coordinate
	Display rune
	G       int
	H       int
	Parent  *Node
}

func (n *Node) F() int {
	return n.G + n.H
}

func (n *Node) IsBlocked() bool {
	return n.Display == Block
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

func (n *Node) CalcGValue(parent *Node) {
	// get distance from parent node
	if abs(parent.Coord.X-n.Coord.X)+abs(parent.Coord.Y-n.Coord.Y) == 2 {
		n.G = DiagonalCost
	} else if abs(parent.Coord.X-n.Coord.X)+abs(parent.Coord.Y-n.Coord.Y) == 1 {
		n.G = DefaultCost
	}
}

func (n *Node) CalcHValue(end *Node) {
	// get distance to end node
	n.H = (abs(end.Coord.X-n.Coord.X) + abs(end.Coord.Y-n.Coord.Y)) * DefaultCost
}
