package main

import (
	"fmt"
	"math/rand"
)

/*
	simple examples are quite idealistic as they didn’t expect any errors happening.
	This is not a good assumption to make in the real world.
	So we should make sure that we handle errors appropriately in production code.
	Luckily there is a package golang.org/x/sync/errgroup that provides a wrapper around the wait group that takes care of the errors.
*/

/*
	Before exploring the error group
	let's first define a job function that can randomly fail.
	We we will make use of it later on.
*/

func Job(jobID int) error {
	if rand.Intn(12) == jobID {
		return fmt.Errorf("job %v failed", jobID)
	}

	fmt.Printf("Job %v done.\n", jobID)
	return nil
}

/*
	This function just generates a random integer between 0 and 11 using the built in rand package.
	If the generated number is equal to the job id then the job fails and returns an error.
	Otherwise it succeeds.
	Now let's look at an example that uses the error group and this job function.
*/
func main() {
	var eg errgroup.Group

	for i := 0; i < 10; i++ {
		jobID := i
		eg.Go(func() error {
			return Job(jobID)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println("Encountered error:", err)
	}
	fmt.Println("Successfully finished.")
}

/*
	This is very similar to how we used the wait group before, maybe even simpler.
	We just create the group and then instead of starting goroutines ourselves and calling Add and Done functions.
	Before starting and when finishing respectively, we just call the Go function on the error group.
	This starts the provided function in a goroutine and then immediately returns, so we can continue executing the loop.
*/
