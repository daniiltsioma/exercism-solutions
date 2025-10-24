package sublist

// Relation type is defined in relations.go file.

// listsEqual checks if two lists are equal.
func listsEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}

	for i, v := range l1 {
		if l2[i] != v {
			return false
		}
	}

	return true
}

func Sublist(l1, l2 []int) Relation {
	// check for equality right away.
	if listsEqual(l1, l2) {
		return "equal"
	} else if len(l1) == len(l2) { 
		// definitely unequal.
		return "unequal"
	}
	
	// check for empty lists.
	if len(l1) == 0 {
		return "sublist"
	}
	if len(l2) == 0 {
		return "superlist"
	}

	// check for sublist.
	if len(l1) < len(l2) {
		for i := 0; i + len(l1) <= len(l2); i++ {
			if listsEqual(l1, l2[i:i+len(l1)]) {
				return "sublist"
			}
		}
	}

	// check for superlist.
	for i := 0; i + len(l2) <= len(l1); i++ {
		if listsEqual(l2, l1[i:i+len(l2)]) {
			return "superlist"
		}
	}

	return "unequal"
}
