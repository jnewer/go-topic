package main

import "fmt"

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

type Number interface {
	int64 | float64
}

func SumNumber[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

type Vector[T any] []T
type NumSlice[T int | float64] []T

type M[K string, V any] map[K]V

type Ch[T any] chan T

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	//fmt.Printf("Generic Sums: %v and %v", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))
	//fmt.Printf("Generic Sums: %v and %v", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n", SumNumber(ints), SumNumber(floats))
	//Generic Sums: 46 and 62.97

	v := Vector[string]{"z", "x", "c"}
	fmt.Printf("v :%v\n", v) //v :[z x c]

	ns := NumSlice[int]{1, 2, 3, 4, 5, 6}
	fmt.Printf("ns :%v\n", ns) //ns :[1 2 3 4 5 6]

	ns2 := NumSlice[float64]{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}
	fmt.Printf("ns2 :%v\n", ns2) //ns2 :[1.1 2.2 3.3 4.4 5.5 6.6]

	m := M[string, int]{
		"zx": 123,
		"as": 456,
		"qw": 789,
	}

	fmt.Printf("m :%v\n", m) //m :map[as:456 qw:789 zx:123]

	ch := make(Ch[int], 1)
	ch <- 10
	res := <-ch

	fmt.Printf("res :%v\n", res) //res :10
	fmt.Printf("ch: %v\n", ch)   //ch: 0xc000010150
}
