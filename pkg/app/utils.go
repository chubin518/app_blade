package app

import (
	"context"
	"errors"
)

// withTimeout
func withTimeout(ctx context.Context, callback func(context.Context) error) error {
	c := make(chan error, 1)
	go func() {
		// If runtime.Goexit() is called from within the callback
		// then nothing is written to the chan.
		// However the defer will still be called, so we can write to the chan,
		// to avoid hanging until the timeout is reached.
		callbackExited := false
		defer func() {
			if !callbackExited {
				// returned when a hook callback does not finish executing
				c <- errors.New("goroutine exited without returning")
			}
		}()
		// returned when callback does not finish executing
		c <- callback(ctx)
		callbackExited = true
	}()

	var err error
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-c:
		// If the context finished at the same time as the callback
		// prefer the context error.
		// This eliminates non-determinism in select-case selection.
		if ctx.Err() != nil {
			err = ctx.Err()
		}
	}
	return err
}
