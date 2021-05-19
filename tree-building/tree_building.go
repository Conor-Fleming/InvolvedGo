package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	fmt.Println(records)
	if len(records) == 0 {
		return nil, nil
	}
	tree := make([]*Node, len(records))

	sort.SliceStable(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	fmt.Println(records)
	for i, v := range records {
		//checking for valid input of records
		if v.ID != i {
			return nil, fmt.Errorf("invalid input, cannot create tree:%+v ", v)
		}
		if v.ID < v.Parent {
			return nil, fmt.Errorf("invalid input, cannot create tree:%+v ", v)
		}
		if i != 0 && v.Parent == i {
			return nil, fmt.Errorf("invalid input, cannot create tree:%+v ", v)
		}

		tree[v.ID] = &Node{ID: v.ID}
		if v.ID != 0 {
			tree[v.Parent].Children = append(tree[v.Parent].Children, tree[v.ID])
		}
	}
	return tree[0], nil
}
