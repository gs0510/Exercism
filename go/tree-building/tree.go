package tree

import (
	"errors"
	"sort"
)

//Record stores the ID and Parent.
type Record struct {
	ID     int
	Parent int
}

//Node stores the ID of the Record and all it's children.
type Node struct {
	ID       int
	Children []*Node
}

//Build builds a tree for given records.
func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	root := &Node{ID: records[0].ID}

	if records[0].ID != 0 || (records[0].ID == 0 && records[0].Parent != 0) {
		return nil, errors.New("the tree has no root node")
	}

	if len(records)-1 != records[len(records)-1].ID {
		return nil, errors.New("the tree is incorrect")
	}

	result := make([]*Node, len(records))

	result[0] = root
	for index := 1; index < len(records); index++ {
		if records[index].Parent >= records[index].ID {
			return nil, errors.New("Bad ID")
		}
		node := Node{ID: records[index].ID}
		result[index] = &node
		parent := result[records[index].Parent]
		parent.Children = append(parent.Children, &node)
	}

	return root, nil

}
