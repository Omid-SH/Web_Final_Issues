/*
	An empty interface can be used to hold any data and it can be a useful parameter since it can work with any type.
	To understand how an empty interface works and how it can hold any type, we should first understand the concept behind the name.

	An interface is two things: it is a set of methods, but it is also a type.
	Interface is composed of two words:
    	• a pointer to information about the type stored
		• a pointer to the associated data
*/

package main

import "fmt"

func main() {

	/*
		var i1 int8 = 1
		// visualize what mainly {} is.
		read1(i1)
	*/

	/*
		var i2 int8 = 1
		//Although the conversion from int8 to a int16 is valid, the program will panic
		read2(i2)
	*/

	/*
		names := []string{"stanley", "david", "oscar"}
		cannot use names (type []string) as type []interface {} in function argument
		PrintAll(names)
	*/

	/*
			If we want to actually make that work,
		    we would have to convert the []string to an []interface{}:
			names := []string{"stanley", "david", "oscar"}
			vals := make([]interface{}, len(names))
			for i, v := range names {
				vals[i] = v
			}
			PrintAll(vals)
	*/

	/*
		dog := Dog{}
		dog.Age = "3"
		fmt.Printf("%#v %T\n", dog.Age, dog.Age) // v -> value, T -> Type.

		dog.Age = 3
		fmt.Printf("%#v %T\n", dog.Age, dog.Age)

		dog.Age = "not really an age"
		fmt.Printf("%#v %T", dog.Age, dog.Age)
	*/

}

type Dog struct {
	Age interface{}
}

func read1(i interface{}) {
	println(i) // prints two addresses -> Both addresses represent the two pointers to "type information" and the "value".
}

// The underlying representation of
// the empty interface is documented in the reflection package:
// type emptyInterface struct {
//	typ  *rtype         // word 1 with type description
//	word unsafe.Pointer // word 2 with the value
// }

func read2(i interface{}) {
	n := i.(int16) // .(destType) type casting in go is
	println(n)
}

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}
