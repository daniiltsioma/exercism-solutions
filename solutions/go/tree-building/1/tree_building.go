package tree

import (
	"errors"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

// visitChildren goes through node's children.
func visitChildren(node *Node, visited map[int]bool) error {
	// node already visited – indirect cycle detected.
	if _, ok := visited[node.ID]; ok {
		return errors.New("cycle indirectly")
	}

	visited[node.ID] = true

	for _, child := range node.Children {
		if err := visitChildren(child, visited); err != nil {
			return err
		}
	}

	return nil
}

// insertChildSorted inserts child sorted by ID.
func insertChildSorted(children []*Node, child *Node) []*Node {
	// if no children – return a new slice.
	if len(children) == 0 {
		return []*Node{child}
	}

	if child.ID < children[0].ID {
		// insert in the beginning.
		return append([]*Node{child}, children...)
	} else if child.ID > children[len(children)-1].ID {
		// insert in the end.
		return append(children, child)
	} else {
		// insert in the middle.
		i := 0
		for children[i].ID < child.ID {
			i++
		}
		return append([](*Node)(children[:i]), append([]*Node{child}, [](*Node)(children[i:])...)...)
	}
}

func Build(records []Record) (*Node, error) {
	// return nil for empty input.
	if len(records) == 0 {
		return nil, nil
	}

	// start by indexing all the records.
	rindex := map[int]*Node{}
	for _, r := range records {
		// check for duplicates.
		if _, exists := rindex[r.ID]; exists {
			return nil, errors.New("duplicate node")
		}
		// check for invalid root.
		if len(records) == 1 && r.Parent > 0 {
			return nil, errors.New("one root node and has parent")
		}
		if r.ID == 0 && r.Parent > 0 {
			return nil, errors.New("root node has parent")
		}
		rindex[r.ID] = &Node{
			ID: r.ID,
			Children: []*Node{},
		}
	}
	
	// link children to parents.
	for i := 0; i < len(records); i++ {
		// check for missing records.
		_, ok := rindex[i]; if !ok {
			return nil, errors.New("non-continuous")
		}

		r := records[i]

		// not linking root to itself.
		if r.ID == 0 {
			continue
		}
		child := rindex[r.ID]
		parent := rindex[r.Parent]

		// check for direct cycles.
		if r.ID != 0 && child == parent {
			return nil, errors.New("cycle directly")
		}

		// check for invalid parent and child IDs.
		if child.ID < parent.ID {
			return nil, errors.New("higher id parent of lower id")
		}

		// insert child in sorted order.
		parent.Children = insertChildSorted(parent.Children, child)
	}

	// check for valid root.
	root, ok := rindex[0]; if !ok {
		return nil, errors.New("no root node")
	}

	// check for indirect cycles or disjoints.
	visited := map[int]bool{}
	if err := visitChildren(root, visited); err != nil {
		return nil, err
	}
	// if not all nodes visited – disjoint(s) detected.
	if len(visited) < len(rindex) {
		return nil, errors.New("non-continuous")
	}

	return root, nil
}
