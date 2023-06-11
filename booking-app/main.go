package main

import "fmt"

func main() {
	s := make([]int, 3, 8)
	fmt.Println(s)

	student := Student{name: "Takeshi", age: 32}
	student.changeName()
	fmt.Println(student)
}

type Student struct {
	name string
	age  int
	_    struct{}
}

func (s *Student) changeName() {
	s.name = "Ito"
}
