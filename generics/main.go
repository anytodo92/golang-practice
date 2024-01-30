package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumInt(p map[string]int64) int64 {
	var s int64
	for _, n := range p {
		s += n
	}

	return s
}

func SumFloat(p map[string]float64) float64 {
	var s float64
	for _, n := range p {
		s += n
	}

	return s
}

func SumIntOrFloat[K comparable, V int64 | float64](p map[K]V) V {
	var s V
	for _, n := range p {
		s += n
	}

	return s
}

func SumNumber[K comparable, V Number](p map[K]V) V {
	var s V
	for _, n := range p {
		s += n
	}

	return s
}

func main() {
	ints := map[string]int64{
		"first":  32,
		"second": 14,
	}

	floats := map[string]float64{
		"first":  1.23,
		"second": 1.35,
	}

	var sumFloat = SumFloat(floats)
	var sumInt = SumInt(ints)

	fmt.Printf("Non-generic Sums: %v and %v\n", sumInt, sumFloat)

	sumFloat = SumIntOrFloat(floats)
	sumInt = SumIntOrFloat[string, int64](ints)

	fmt.Printf("Generic Sums: %v and %v\n", sumInt, sumFloat)

	sumFloat = SumNumber(floats)
	sumInt = SumNumber(ints)
	fmt.Printf("Constraint Sums: %v and %v\n", sumInt, sumFloat)
}
