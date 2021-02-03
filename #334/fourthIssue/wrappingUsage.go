// whats the use case of wrapping the errors.
// To understand it let’s see an example

package main

import (
	"fmt"
)

type notPositive struct {
	num int
}

func (e notPositive) Error() string {
	return fmt.Sprintf("checkPositive: Given number %d is not a positive number", e.num)
}

type notEven struct {
	num int
}

func (e notEven) Error() string {
	return fmt.Sprintf("checkEven: Given number %d is not an even number", e.num)
}

func checkPositive(num int) error {
	if num < 0 {
		return notPositive{num: num}
	}
	return nil
}

func checkEven(num int) error {
	if num%2 == 1 {
		return notEven{num: num}
	}
	return nil
}

func checkPostiveAndEven(num int) error {
	if num > 100 {
		return fmt.Errorf("checkPostiveAndEven: Number %d is greater than 100", num)
	}

	err := checkPositive(num)
	if err != nil {
		return err
	}

	err = checkEven(num)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	num := 3
	err := checkPostiveAndEven(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Givennnumber is positive and even")
	}

}
