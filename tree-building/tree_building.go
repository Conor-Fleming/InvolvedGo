//Tree package lets us build nice trees to dispplay
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

//Build function takes a slice of Record objects and creates a tree based on ID and Parent values of each record
func Build(records []Record) (*Node, error) {
	fmt.Println(records)
	if len(records) == 0 {
		return nil, nil
	}
	tree := make([]*Node, len(records))

	sort.SliceStable(records, func(i, j int) bool { return records[i].ID < records[j].ID })
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

		//create Node and inserting in tree
		//if the id of the current Record is NOT 0(not root of tree) then we append this node to its parents Children array
		tree[v.ID] = &Node{ID: v.ID}
		if v.ID != 0 {
			tree[v.Parent].Children = append(tree[v.Parent].Children, tree[v.ID])
		}
	}
	//return the root of the created tree
	return tree[0], nil
}
