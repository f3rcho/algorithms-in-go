package patterns

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// someTask function that we call periodically.
func someTask() {
	fmt.Println(rand.Int() * rand.Int())
}

func periodicTask(ctx context.Context) {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			someTask()
		case <-ctx.Done():
			fmt.Println("stopping PeriodicTask")
			ticker.Stop()
			return
		}
	}
}

func RunSelect() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go periodicTask(ctx)

	// Create a channel to receive signals from the operating system.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)

	// The code blocks until a signal is received (e.g. Ctrl+C).
	<-sigCh
}

// When to Use This Pattern
// This pattern is useful when you need to perform a task in an infinite loop based on some event or timer and then stop its execution based on a certain condition.

// For example, it can be used to run deferred calculations using data that was saved to a database, or to asynchronously enrich records in the database with data from other services. At the same time, we always have the ability to safely terminate the goroutine when the context is canceled or some other external event occurs.
