package school

import (
	"sort"
)

type School struct {
	grades []Grade
}

type Grade struct {
	grade    int
	students GradeSlice
}

// ###################################################
// For implementing Sort - we need interfaces
// See Page 187 in book PDF
// School
func (s School) Len() int { return len(s.grades) }
func (s School) Less(i, j int) bool { 
	return s.grades[i].grade < s.grades[j].grade 
}
func (s School) Swap(i, j int) { s.grades[i], s.grades[j] = s.grades[j], s.grades[i]}
// Grade
type GradeSlice []string
func (g GradeSlice) Len() int {return len(g)}
func (g GradeSlice) Less(i, j int) bool { 
	return g[i][0] < g[j][0] // First letter of name
}
func (g GradeSlice) Swap(i, j int) { g[i], g[j] = g[j], g[i]}
// ###################################################

// Create and return a new, empty school
func New() *School {
	school := School{[]Grade{}}
	return &school
}

// Add student to a Grade - whether or not the Grade exists is handled
func (s *School) Add(student string, g int) {
	if !s.HasGrade(g) { // If Grade doesn't exist, add to school with student appended
		s.grades = append(s.grades, Grade{grade: g, students: []string{student}})
	} else { // Grade already exists, append student to grade
		s.AppendStudentToGrade(g, student)
	}
}

// For an existing Grade in the school, append student to it
func (s *School) AppendStudentToGrade(level int, student string) {
	for i, g := range(s.grades) {
		if g.grade == level {
			s.grades[i].students = append(s.grades[i].students, student)
		}
	}
}

// Return the students in a Grade
func (s *School) Grade(level int) []string {
	for _, g := range(s.grades) {
		if g.grade == level {
			return g.students
		}
	}
	return []string{}
}

// Return grade enrollments, sorted in place
func (s *School) Enrollment() []Grade {
	sort.Sort(s)
	for _, g := range(s.grades) {
		sort.Sort(g.students)
	}
	return s.grades
}

// Whether or not school has students in a Grade enrolled
func (s *School) HasGrade(grade int) bool {
	grades := s.grades
	for _, g := range grades {
		if g.grade == grade {
			return true
		}
	}
	return false
}
