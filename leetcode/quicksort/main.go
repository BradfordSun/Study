package main

import (
	"fmt"
)

// 找到基线pivot（数组第一位），生成比他大的数组及比他小的数组，把小的，基线和大的拼一块
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]
		arr := arr[1:]
		less := func() []int {
			var l []int
			for i, _ := range arr {
				if arr[i] <= pivot {
					l = append(l, arr[i])
				}
			}
			return l
		}
		greater := func() []int {
			var g []int
			for i, _ := range arr {
				if arr[i] > pivot {
					g = append(g, arr[i])
				}
			}
			return g
		}
		result := append(quickSort(less()), pivot)
		result = append(result, quickSort(greater())...)
		return result
	}
}

func main() {
	a := []int{8, 102, 2, 3, 90, 5, 0, 100}
	result := quickSort(a)
	fmt.Println(result)
}
