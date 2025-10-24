package stringset

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]string

func New() Set {
	return map[string]string{}
}

func NewFromSlice(l []string) Set {
	m := map[string]string{}
	for _, s := range l {
		m[s] = s
	}
	return m
}

func (s Set) String() string {
	str := "{"
	i := 0
	for m := range s {
		if i > 0 {
			str += ", "
		}
		str += "\""	+ m + "\""
		i++
	}
	str += "}"

	return str
}

func (s Set) IsEmpty() bool {
	for _ = range s {
		return false
	}
	return true
}

func (s Set) Has(elem string) bool {
	for m := range s {
		if elem == m {
			return true
		}
	}
	return false
}

func (s Set) Add(elem string) {
	s[elem] = elem
}

func Subset(s1, s2 Set) bool {
	for m1 := range s1 {
		_, ok := s2[m1]; if !ok {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for m1 := range s1 {
		_, ok := s2[m1]; if ok {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	s := Set{}
	for m1 := range s1 {
		_, ok := s2[m1]; if ok {
			s.Add(m1)
		}
	}

	return s
}

func Difference(s1, s2 Set) Set {
	s := Set{}
	for m1 := range s1 {
		_, ok := s2[m1]; if !ok {
			s.Add(m1)
		}
	}

	return s
}

func Union(s1, s2 Set) Set {
	s := Set{}
	for m1 := range s1 {
		s.Add(m1)
	}
	for m2 := range s2 {
		s.Add(m2)
	}

	return s
}
