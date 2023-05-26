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
	r2 := Reduce(arr, func(res *[]int, item int, k int) {
		*res = append(*res, item*item)
	})
	fmt.Println(r2)
	type Tage struct {
		age int
	}
	r3 := Reduce(arr, func(res *Tage, item int, k int) {
		res.age += item
	})
	fmt.Println(r3)
}
