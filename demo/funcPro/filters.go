package funcPro

func Filter(students []*Student, f func(*Student) bool) []*Student {
	result := []*Student{}
	for _, student := range students {
		if f(student) {
			result = append(result, student)
		}
	}
	return result
}

func AgeGreatThanFunc(age int) func(*Student) bool {
	return func(s *Student) bool {
		return s.Age > age
	}
}

func ComplexFunc(age int, height int) func(*Student) bool {
	return func(s *Student) bool {
		return (s.Age > age) && (s.Height > height)
	}
}
