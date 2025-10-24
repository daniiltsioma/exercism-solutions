package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	res := map[string]int{}

	for k := range in {
		for _, v := range in[k] {
			res[strings.ToLower(v)] = k
		}
	}		
	
	return res
}
