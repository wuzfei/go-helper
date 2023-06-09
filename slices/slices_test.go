package slices

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9}
	r1 := Reduce(arr, func(res *int, item int, k int) {
		*res = *res + item
	})
	fmt.Println(r1)
	r2 := Reduce(arr, func(res *map[int]int, item int, k int) {
		if *res == nil {
			*res = make(map[int]int)
		}
		(*res)[item] = item
	})
	fmt.Println(r2)
	r3 := Reduce(arr, func(res *[]int, item int, k int) {
		*res = append(*res, item*item)
	})
	fmt.Println(r3)
	type Tage struct {
		age int
	}
	r4 := Reduce(arr, func(res *Tage, item int, k int) {
		res.age += item
	})
	fmt.Println(r4)
}
