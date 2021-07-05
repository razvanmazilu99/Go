package entity

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (person Person) IsMale() bool {
	if person.Gender == "Male" {
		return true
	}

	return false
}