package funcs

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRunTimer(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	err := RunTimer(ctx, time.Second*5, func() error {
		fmt.Println("timeout")
		return nil
	})

	fmt.Println("1111", err)

}
func TestRunTicker(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second*7)

	RunTicker(ctx, time.Second*2, func() error {
		fmt.Println("timeout")
		return nil
	})

	fmt.Println("1111")

}
