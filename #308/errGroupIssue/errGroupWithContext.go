package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
	The errgroup package provides an alternative way to create an error group using context by calling the WithContext function.
	This function takes a context as an argument and returns an error group along with a context that is a child of the provided one
	(this means that whenever the parent context is cancelled, the child one is cancelled as well).
*/

/*
	The context returned by the WithContext function is also cancelled once
	any of the functions started in the error group terminates with an error.
	This will allow us to stop the execution of the other jobs in the group
	as soon as we encounter the first error.
	(Note that the error group context is also cancelled once the first call to Wait returns as well, so the context can't be reused.)
*/
func JobWithCtx(ctx context.Context, jobID int) error {
	select {
	case <-ctx.Done():
		fmt.Printf("context cancelled job %v terminting\n", jobID)
		return nil
	case <-time.After(time.Second * time.Duration(rand.Intn(3))):
	}
	if rand.Intn(12) == jobID {
		fmt.Printf("Job %v failed.\n", jobID)
		return fmt.Errorf("job %v failed", jobID)
	}

	fmt.Printf("Job %v done.\n", jobID)
	return nil
}

func main() {
	eg, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < 10; i++ {
		jobID := i
		eg.Go(func() error {
			return JobWithCtx(ctx, jobID)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println("Encountered error:", err)
	}
	fmt.Println("Successfully finished.")
}
