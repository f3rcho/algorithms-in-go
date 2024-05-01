package patterns

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// errFailure some custom error.
var errFail = errors.New("my bad err")

func RunErrGrCtx() {
	// Create errgroup with context.
	group, qctx := errgroup.WithContext(context.Background())

	// Run first periodic task.
	group.Go(func() error {
		firstTask(qctx)
		return nil
	})

	// run second task.
	group.Go(func() error {
		if err := secondTask(); err != nil {
			return err
		}
		return nil
	})
	// Wait for all tasks to complete or the error to appear.
	if err := group.Wait(); err != nil {
		fmt.Printf("errGroup tasks ended up with an error: %v", err)
	}
}

func firstTask(ctx context.Context) {
	var counter int
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("some task...")
			if counter > 10 {
				return
			}
			counter++
		}
	}
}

func secondTask() error {
	time.Sleep(3 * time.Second)
	return errFail
}

/*
When to Use This Pattern

I use this pattern where it is important for all goroutines in the group to complete successfully without errors.

For example, if I need to perform a calculation using combined data. I cannot perform the calculation if data from a particular table is missing (e.g., it has not yet been saved). In this case, I return a custom error about the missing data, and the execution of all other database queries is interrupted until the next task group run.
*/
