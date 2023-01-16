package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

type ByID []Record
func (r ByID) Len() int {return len(r)}
func (r ByID) Swap(i, j int) {r[i], r[j] = r[j], r[i]}
func (r ByID) Less(i, j int) bool {return r[i].ID < r[j].ID}

func Build(records []Record) (*Node, error) {
	sort.Sort(ByID(records))

	nodes := make([]*Node, len(records))

	if len(records) == 0 {
		return nil, nil
	}
	if records[0].ID != 0 || records[0].Parent != 0 {
		return nil, fmt.Errorf("No root node")
	}

	for i, r := range records {
		nodes[i] = &Node{ID: r.ID}

		if i == 0 {
			continue
		} else if i != r.ID || r.ID <= r.Parent {
			return nil, fmt.Errorf("Not a valid record")
		}

		if parent := nodes[r.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[i])
		}
	}

	return nodes[0], nil
}