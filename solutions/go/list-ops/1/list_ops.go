package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := s.Length() - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	res := make(IntList, s.Length())
	
	i := 0
	for _, v := range s {
		if fn(v) {
			res[i] = v
			i++
		}
	}

	return res[:i]
}

func (s IntList) Length() int {
	l := 0
	for _ = range s {
		l++
	}
	return l
}

func (s IntList) Map(fn func(int) int) IntList {
	res := make(IntList, s.Length())
	for i, v := range s {
		res[i] = fn(v)
	}
	return res
}

func (s IntList) Reverse() IntList {
	res := make(IntList, s.Length())
	for i, j := s.Length() - 1, 0; i >= 0; i, j = i-1, j+1 {
		res[i] = s[j]
	}
	return res
}

func (s IntList) Append(lst IntList) IntList {
	newList := make(IntList, s.Length() + lst.Length())
	for i, v := range s {
		newList[i] = v
	}
	
	for i, j := s.Length(), 0; i < newList.Length(); i, j = i+1, j+1 {
		newList[i] = lst[j]
	}
	return newList
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, l := range lists {
		s = s.Append(l)
	}

	return s
}
