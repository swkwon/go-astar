package astar

import (
	"fmt"
	"slices"
)

type MapInfo struct {
	Tag    string
	Start  *Node
	End    *Node
	Width  int
	Height int
	Map    map[int]map[int]*Node

	OpenSet   *NodeSet
	ClosedSet *NodeSet
}

func (m *MapInfo) Print() {
	fmt.Printf("\n%s\n", m.Tag)
	for i := 0; i < m.Width; i++ {
		for j := 0; j < m.Height; j++ {
			fmt.Printf("%c ", m.Map[i][j].Display)
		}

		fmt.Println()
	}
}

func (m *MapInfo) GetNearNodes(node *Node) []*Node {
	var ret []*Node
	search := []struct{ x, y int }{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	for _, v := range search {
		nearX := node.Coord.X + v.x
		nearY := node.Coord.Y + v.y
		if nearX < 0 || nearX >= m.Width || nearY < 0 || nearY >= m.Height {
			continue
		}
		nearNode := m.Map[nearX][nearY]
		if m.OpenSet.IsInSet(&nearNode.Coord) ||
			m.ClosedSet.IsInSet(&nearNode.Coord) ||
			nearNode.IsBlocked() {
			continue
		}
		ret = append(ret, nearNode)
	}
	return ret
}

func (m *MapInfo) Find() []*Node {
	var ret []*Node

	curNode := m.Start
	m.OpenSet.Push(curNode)

	for m.OpenSet.Len() > 0 {
		m.ClosedSet.Push(curNode)
		m.OpenSet.Pop()

		if curNode == m.End {
			break
		}

		nearNodes := m.GetNearNodes(curNode)
		for _, node := range nearNodes {
			node.Parent = curNode
			node.CalcGValue(curNode)
			node.CalcHValue(m.End)
			m.OpenSet.Push(node)
		}

		// find the node with the smallest f value in the open set.
		smallestNode := m.OpenSet.GetRoot()
		if smallestNode == nil {
			break
		} else {
			curNode = smallestNode
		}
	}

	if m.End.Parent != nil {
		var n *Node
		n = m.End
		for n != nil {
			if n != m.Start && n != m.End {
				n.Display = Dot
			}

			ret = append(ret, n)
			n = n.Parent
		}
	}

	slices.Reverse(ret)
	return ret
}
