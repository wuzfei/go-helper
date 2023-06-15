package funcs

import (
	"context"
	"time"
)

// RunTimer 定时器执行一次
func RunTimer(ctx context.Context, t time.Duration, fn func() error) error {
	tt := time.NewTimer(t)
	defer tt.Stop()
	select {
	case <-tt.C:
		return fn()
	case <-ctx.Done():
	}
	return nil
}

// RunTicker 定时器间隔执行
func RunTicker(ctx context.Context, t time.Duration, fn func() error) {
	tk := time.NewTicker(t)
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			if fn() != nil {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}
