package school

import (
	"sort"
	"slices"
)

// Define the Grade and School types here.
type Grade struct {
	level int
	students []string
}

type School struct {
	grades map[int]*Grade
}

func New() *School {
	return &School{
		grades: map[int]*Grade{},
	}
}

func (s *School) Add(student string, grade int) {
	g := s.grades[grade]
	if g != nil {
		g.students = append(g.students, student)
	} else {
		s.grades[grade] = &Grade{
			level: grade,
			students: []string{student},
		}
	}
}

func (s *School) Grade(level int) []string {
	g, ok := s.grades[level]; if !ok {
		return []string{}
	}
	return g.students
}

func (s *School) Enrollment() []Grade {
	grades := []Grade{}
	for _, g := range s.grades {
		sort.Strings(g.students)
		grades = append(grades, *g)
	}
	slices.SortFunc(grades, func(g1, g2 Grade) int {
		return g1.level - g2.level
	})

	return grades
}
