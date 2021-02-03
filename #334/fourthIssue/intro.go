/*

introducing
1-wrapping,unwrapping(for usage see wrappingUsage.go)
2-error.Is
3-error.As

*/

package main

import (
	"errors"
	"fmt"
)

type errorOne struct{}

func (e errorOne) Error() string {
	return "Error One happended"
}

func main() {

	/*
		In go, error can wrap another error as well.
		What does the wrapping of error mean?
		It means to create a hierarchy of errors in which a  particular instance of
		error wraps another error and that particular instance itself can be
		wrapped inside another error.  Below is the syntax for wrapping an error
		e := fmt.Errorf("... %w ...", ..., err, ...)

	*/

	e1 := errorOne{}

	e2 := fmt.Errorf("E2: %w", e1)

	e3 := fmt.Errorf("E3: %w", e2)

	fmt.Println(e2)

	fmt.Println(e3)

	/*
		In the above section, we studied about wrapping the error.
		It is also possible to unwrap the error.
		Unwrap function of errors package can be used to unwrap an error.
		Below is the syntax of the function.
		func Unwrap(err error) error
	*/

	e1 := errorOne{}
	e2 := fmt.Errorf("E2: %w", e1)
	e3 := fmt.Errorf("E3: %w", e2)
	fmt.Println(errors.Unwrap(e3))
	fmt.Println(errors.Unwrap(e2))
	fmt.Println(errors.Unwrap(e1))

	/*
		Is unwraps its first argument sequentially looking for an error that matches the second. It reports whether it finds a match. It should be used in preference to simple equality checks:

		if errors.Is(err, os.ErrExist)
		is preferable to

		if err == os.ErrExist
	*/

	/*
		As unwraps its first argument sequentially looking for
		an error that can be assigned to its second argument,
		which must be a pointer. If it succeeds, it performs the assignment
		and returns true. Otherwise, it returns false. The form

		var perr *os.PathError
		if errors.As(err, &perr) {
			fmt.Println(perr.Path)
		}
		is preferable to

		if perr, ok := err.(*os.PathError); ok {
			fmt.Println(perr.Path)
		}
	*/

}
