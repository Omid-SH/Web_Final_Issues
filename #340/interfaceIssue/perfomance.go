/*
	Any conversion from the inner type of an empty interface should be done
	after the conversion of the original type.
	This conversion to an empty interface and then to the original type back has
	a cost for your program.
	(we have some benchmarks to evaluate this).
*/

package main

import (
	"testing"
)

type MultipleFieldStructure struct {
	a int
	b string
	c float32
	d float64
	e int32
	f bool
	g uint64
	h *string
	i uint16
}

var x MultipleFieldStructure

/*
	"low perfomance and speed."
	SimpleTest:
	BenchmarkWithType-8 300000000 4.24 ns/op
	BenchmarkWithEmptyInterface-8 20000000 60.4 ns/op
*/
func emptyInterface(i interface{}) {
	s := i.(MultipleFieldStructure)
	x = s
}

var y *MultipleFieldStructure

/*
	"high perfomance and speed."
	A good solution would be to use pointer and convert back to this same struct pointer.
	Results with the same test:
	BenchmarkWithType-8 2000000000 2.16 ns/op
	BenchmarkWithEmptyInterface-8 2000000000 2.02 ns/op // perfomance enhanced.
*/
func emptyInterface_enhanced(i interface{}) {
	s := i.(*MultipleFieldStructure)
	y = s
}

func typed(s MultipleFieldStructure) {
	x = s
}

func BenchmarkWithType(b *testing.B) {
	s := MultipleFieldStructure{a: 1, h: new(string)}
	for i := 0; i < b.N; i++ {
		typed(s)
	}
}

func BenchmarkWithEmptyInterface(b *testing.B) {
	s := MultipleFieldStructure{a: 1, h: new(string)}
	for i := 0; i < b.N; i++ {
		emptyInterface(s)
	}
}
