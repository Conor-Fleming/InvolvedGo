package tree

type Record struct {
	ID     int
	Parent int
}
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	tree := make([]*Node, 0)
	root := Node
	for records != nil {
		for i, v := range records {
			if v.ID == 0 {
				root := Node{0, nil}
				tree[i] = &root

			}
		}
	}

}
