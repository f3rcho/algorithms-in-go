package patterns

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

var errFailure = errors.New("my bad")

func RunErrGroup() {
	group := errgroup.Group{}

	group.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("doing some work 1")
		return nil
	})

	// run second task
	group.Go(func() error {
		fmt.Println("doing some taks 2")
		return nil
	})

	// run third
	group.Go(func() error {
		fmt.Println("doing some work 3")
		return errFailure
		// return nil
	})

	//wait for all goroutines to complete
	if err := group.Wait(); err != nil {
		fmt.Printf("errgroup taks up with an error:%v\n", err)
	} else {
		fmt.Println("all workds done successfully")
	}
}
