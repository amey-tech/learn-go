package methods

import (
	"fmt"
	"math"
)

// Both Scale and Abs are methods with receiver type *Coordinate
// Even though the Abs method needn't modify its receiver

type Coordinate struct {
	X, Y float64
}

func (v *Coordinate) Abs() float64 {
	if v == nil {
		fmt.Println("<nil>")
		return float64(0)
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Coordinate) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.

type Absoluteness interface {
	Abs() float64
}

type AbsolutenessByValue interface {
	Abs() float64
}

// A type implements an interface by implementing its methods.
// There is no explicit declaration of intent, no "implements" keyword.
// Implicit interfaces decouple the definition of an interface from its implementation,
// which could then appear in any package without prearrangement.

type MyFloat float64

// This method means type MyFloat implements the interface Absoluteness,
// but we don't need to explicitly declare that it does so.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
// (value, type)
// An interface value holds a value of a specific underlying concrete type.
// Calling a method on an interface value executes the method of the same name on its underlying type.
// Below is a function to show value and type information
func Describe(i Absoluteness) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func DescribeGeneric(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func DemoImplementationMethodsAndInterface() {
	var a Absoluteness

	myFloat := MyFloat(-math.Sqrt(2))
	myCoordinate := Coordinate{-3, -4}

	a = myFloat // a MyFloat implements Absoluteness
	fmt.Println("Abs method called on MyFloat:", a.Abs())
	Describe(a)
	// DescribeGeneric(a)

	a = &myCoordinate // a *Coordinate implements Absoluteness
	fmt.Println("Abs method called on Coordinate:", a.Abs())
	Describe(a)
	// DescribeGeneric(a)

	// In the following line, myCoordinate is a Coordinate (not *Coordinate) and does NOT implement Absoluteness.
	// a = myCoordinate

	// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
	// In some languages this would trigger a null pointer exception, but in Go it is common to write
	// methods that gracefully handle being called with a nil receiver (as with the method M in this example.)
	// Note that an interface value that holds a nil concrete value is itself non-nil.
	var b *Coordinate
	Describe(b)
	// DescribeGeneric(b)
	b.Abs()

	// A nil interface value holds neither value nor concrete type.
	// Calling a method on a nil interface is a run-time error because
	// there is no type inside the interface tuple to indicate which concrete method to call.
	// Check by uncommenting the following lines
	// var c Absoluteness
	// Describe(c)
	// c.Abs()

	// The interface type that specifies zero methods is known as the empty interface: interface{}
	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	// Empty interfaces are used by code that handles values of unknown type.
	// Eg: fmt.Print takes any number of arguments of type interface{}.
	var i interface{}
	DescribeGeneric(i)
	i = 42
	DescribeGeneric(i)
	i = "hello"
	DescribeGeneric(i)
}
