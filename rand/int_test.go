package rand

import (
	"fmt"
	"testing"
)

func TestBetweenN(t *testing.T) {
	n1 := 10
	n2 := 20
	fmt.Println(BetweenN(n1, n2, 3))
}

func TestBetween(t *testing.T) {
	n1 := 10
	n2 := 20
	fmt.Println(Between(n1, n2))
}
