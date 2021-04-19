package main

import "fmt"

const (
	Student = iota
	_
	Teacher = "a"
	Leader = iota
)

func main() {
	fmt.Println(Student,Teacher,Leader)
}
