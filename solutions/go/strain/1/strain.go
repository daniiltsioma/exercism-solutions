package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](collection []T, predicate func (T) bool) []T {
	var result = []T{}

	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func Discard[T any](collection []T, predicate func (T) bool) []T {
	var result = []T{}

	for _, item := range collection {
		if !predicate(item) {
			result = append(result, item)
		}
	}

	return result
}