package main

import (
	"fmt"
	"math"
)

const Delta = 0.0001

func isConverged(d float64) bool {
	if d < 0.0 {
		d = -d
	}
	if d < Delta {
		return true
	}
	return false
}

func Sqrt(x float64) float64 {
	z := 1.0
	tmp := 0.0
	for {
		tmp = z - (z*z-x)/2*z
		if d := tmp - z; isConverged(d) {
			return tmp
		}
		z = tmp
	}
	return z
}

type Student struct {
	name    string
	address string
}

// Structs
func TestStruct() {
	fmt.Println(Student{"Sun*", "HaNoi"})
}

// Arrays
func TestArray() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// Slices
func TestSlices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[0:4]
	fmt.Println(s)
}

// Maps
func TestMaps() {
	m := make(map[string]Student)
	fmt.Println(m)

	m["Bell Labs"] = Student{
		"Sun", "Asterisk",
	}
	fmt.Println(m["Bell Labs"])
}

type Vertex struct {
	X, Y float64
}

// Methods
func (v Vertex) TestMethods() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// For loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	TestStruct()
	TestArray()
	TestSlices()
	TestMaps()

	v := Vertex{3, 4}
	fmt.Println(v.TestMethods())
}
