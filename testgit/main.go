package main

import "fmt"

func add(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 10
	b := 20
	add(&a, &b)
	fmt.Println(a, b)
}
