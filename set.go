package astar

import (
	"errors"
)

type SetType int

const (
	Open   SetType = 1
	Closed SetType = 2
)

type NodeSet struct {
	Type         SetType
	Nodes        []*Node
	NodeIndexMap map[string]int
}

func MakeNodeSet(setType SetType) (*NodeSet, error) {

	switch setType {
	case Open:
		return &NodeSet{
			Type:         setType,
			NodeIndexMap: make(map[string]int),
		}, nil
	case Closed:
		return &NodeSet{
			Type:         setType,
			NodeIndexMap: make(map[string]int),
		}, nil
	default:
		return nil, errors.New("invalid set type")
	}
}

func (ns *NodeSet) Push(node *Node) {
	ns.Nodes = append(ns.Nodes, node)
	ci := len(ns.Nodes) - 1
	ns.NodeIndexMap[node.Coord.String()] = ci

	if ns.Type == Open {
		// bubble up
		for ci != 0 {
			pi := (ci - 1) / 2
			if ns.Nodes[pi].F() > ns.Nodes[ci].F() {
				temp := ns.Nodes[pi]
				ns.Nodes[pi] = ns.Nodes[ci]
				ns.NodeIndexMap[ns.Nodes[pi].Coord.String()] = pi
				ns.Nodes[ci] = temp
				ns.NodeIndexMap[ns.Nodes[ci].Coord.String()] = ci
				ci = pi
			} else {
				break
			}
		}
	}
}

func (ns *NodeSet) Pop() {
	if len(ns.Nodes) <= 0 {
		return
	}
	pi := 0
	node := ns.Nodes[pi]

	delete(ns.NodeIndexMap, node.Coord.String())
	lastIndex := len(ns.Nodes) - 1
	ns.Nodes[pi] = ns.Nodes[lastIndex]
	ns.Nodes = ns.Nodes[:lastIndex]
	lastIndex = lastIndex - 1

	if ns.Type == Open {
		// sink down
		for {
			left, right := (pi*2)+1, (pi*2)+2
			if left > lastIndex {
				break
			}

			// 기본적으로 왼쪽 child로 비교 설정
			ci := left
			// 오른쪽이 작으면 오른쪽 child와 비교하도록 설정
			if right <= lastIndex && ns.Nodes[left].F() > ns.Nodes[right].F() {
				ci = right
			}
			// parent와 child의 f값 비교 시 parent가 작거나 같으면 loop 탈출
			if ns.Nodes[pi].F() <= ns.Nodes[ci].F() {
				break
			}
			// swap
			temp := ns.Nodes[pi]
			ns.Nodes[pi] = ns.Nodes[ci]
			ns.NodeIndexMap[ns.Nodes[pi].Coord.String()] = pi
			ns.Nodes[ci] = temp
			ns.NodeIndexMap[ns.Nodes[ci].Coord.String()] = ci
			pi = ci
		}
	}
}

func (ns *NodeSet) IsInSet(coord *Coordinate) bool {
	_, ok := ns.NodeIndexMap[coord.String()]
	return ok
}

func (ns *NodeSet) GetRoot() *Node {
	if len(ns.Nodes) > 0 {
		return ns.Nodes[0]
	}
	return nil
}

func (ns *NodeSet) Len() int {
	return len(ns.Nodes)
}
