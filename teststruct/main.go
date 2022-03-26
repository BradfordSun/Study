package main

import "fmt"

type Student struct {
	Name string
	age  int
}

type JustTest interface {
	PrintName()
	ChangeAge()
}

func (s *Student) PrintName() {
	fmt.Println(s.Name)
}

func (s *Student) ChangeAge() {
	s.age++
	fmt.Println(s.age)
}

func main() {
	ztt := &Student{
		Name: "zhangtingting",
		age:  18,
	}
	var justtest JustTest = ztt
	justtest.ChangeAge()
	fmt.Println(*ztt)
	justtest.PrintName()
}
