package main

import (
	"fmt"
)

type Student struct {
	Name string
	age  int
}

func (s *Student) WorkerChan() chan int {
	return make(chan int)
}

type JustTest interface {
	PrintName()
	ChangeAge()
	WorkerChan() chan int
}

func (s *Student) PrintName() {
	fmt.Println(s.Name)
}

func (s *Student) ChangeAge() {
	s.age++
	fmt.Println(s.age)
}

func test(in chan int) {
	aaa := <-in
	fmt.Println(aaa)
}

func main() {
	ztt := &Student{
		Name: "zhangtingting",
		age:  18,
	}
	var justtest JustTest = ztt
	justtest.ChangeAge()
	justtest.PrintName()
	test(ztt.WorkerChan())
}
