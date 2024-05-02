package astar

const (
	Available = '0' // Available
	Block     = 'X' // Block
	Start     = 'S' // Start
	End       = 'E' // End
	Dot       = '.' // Path
)

const (
	DefaultCost  = 10
	DiagonalCost = 14
)

func MakeDefaultMap(tag string, width, height int, start, end Coordinate, block []Coordinate) (*MapInfo, error) {
	ret := &MapInfo{}

	ret.Map = make(map[int]map[int]*Node)
	ret.Tag = tag
	ret.Width = width
	ret.Height = height
	var err error
	ret.OpenSet, err = MakeNodeSet(Open)
	if err != nil {
		return nil, err
	}

	ret.ClosedSet, err = MakeNodeSet(Closed)
	if err != nil {
		return nil, err
	}

	for i := 0; i < width; i++ {
		if ret.Map[i] == nil {
			ret.Map[i] = make(map[int]*Node)
		}
		for j := 0; j < height; j++ {
			node := &Node{
				Coord:   Coordinate{i, j},
				Display: getDisplay(i, j, start, end, block),
			}

			if start.X == i && start.Y == j {
				ret.Start = node
			}

			if end.X == i && end.Y == j {
				ret.End = node
			}

			ret.Map[i][j] = node
		}
	}

	return ret, nil
}

func getDisplay(i, j int, start, end Coordinate, block []Coordinate) rune {
	if start.X == i && start.Y == j {
		return Start
	} else if end.X == i && end.Y == j {
		return End
	}
	for _, v := range block {
		if v.X == i && v.Y == j {
			return Block
		}
	}
	return Available
}
