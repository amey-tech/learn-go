package methods

import (
	"fmt"
	"math"
)

// Go does not have classes. However, we can define methods on types.
// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// The following method has a receiver of type Vertex named v.

type Vertex struct {
	X, Y float64
}

func (v Vertex) Absolute() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// A method is just a function with a receiver argument.
// Here's a regular function with no change in functionality.

func AbsoluteFunction(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// We can declare a method on non-struct types, too.
// In this example we see a numeric type MyCustomFloat with an Abs method.
// We can only declare a method with a receiver whose type is defined in the same package as the method.
// We cannot declare a method with a receiver whose type is defined in another package
// (which includes the built-in types such as int).

type MyCustomFloat float64

func (f MyCustomFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Try doing this ->
// func (f float64) TryAbs() float64 {
// 	if f < 0 {
// 		return -f
// 	}
// 	return f
// }

// We can declare methods with pointer receivers.
// This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

// For example, the Scale method here is defined on *Vertex.

// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
// Try removing the * from the declaration of the Scale function.

// With a value receiver, the Scale method operates on a copy of the original Vertex value.
// (This is the same behavior as for any other function argument.)
// The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

func (v Vertex) ScaleWithValue(f float64) {
	// Co-pilot says -
	// To solve the problem of ineffective assignment to field Vertex.X
	// we need to use a pointer receiver for the ScaleWithValue method.
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) ScaleWithPointer(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// The above methods as functions
func ScaleWithValueFunction(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleWithPointerFunction(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func DemoImplementationMethodsIntroduction() {
	v1 := Vertex{X: 3, Y: 4}
	fmt.Println("Method call (v1):", v1.Absolute())
	fmt.Println("Function call (v1):", AbsoluteFunction(v1))

	myCustomFloat := MyCustomFloat(-10)
	fmt.Println("Abs method call (v1):", myCustomFloat.Abs())

	v1.ScaleWithValue(10)
	fmt.Println("Value receiver method call (v1):", v1, v1.Absolute())
	v1.ScaleWithPointer(10)
	fmt.Println("Pointer receiver method call (v1):", v1, v1.Absolute())

	//Reset v1
	v1 = Vertex{X: 3, Y: 4}
	ScaleWithValueFunction(v1, 10)
	fmt.Println("Function call with value (v1):", v1, v1.Absolute())
	ScaleWithPointerFunction(&v1, 10)
	fmt.Println("Function call with pointer (v1):", v1, v1.Absolute())

	// We noticed that functions with a pointer argument must take a pointer:
	v2 := Vertex{X: 3, Y: 4}
	p2 := &v2

	// var v Vertex
	// ScaleWithPointerFunction(v, 2)  -> Compile error!
	// ScaleWithPointerFunction(&v, 2) -> OK
	ScaleWithPointerFunction(&v2, 2)
	fmt.Println("Function call with pointer (v2):", v2)
	ScaleWithPointerFunction(p2, 2)
	fmt.Println("Function call with pointer (p2):", v2)

	// While methods with pointer receivers take either a value or a pointer as the receiver when they are called:
	// var v Vertex
	// v.ScaleWithPointer(2)  -> OK
	// p := &v
	// p.ScaleWithPointer(2) -> OK
	v2.ScaleWithPointer(2) // (&v2).ScaleWithPointer(2)
	fmt.Println("Pointer receiver method call (v2):", v2)
	p2.ScaleWithPointer(2)
	fmt.Println("Pointer receiver method call (p2):", v2)
	// For the statement v.ScaleWithPointer(2), even though v is a value and not a pointer
	// the method with the pointer receiver is called automatically
	// That is, as a convenience, Go interprets the statement v.ScaleWithPointer(2) as (&v).ScaleWithPointer(2)
	// since the ScaleWithPointer method has a pointer receiver

	// The equivalent thing happens in the reverse direction.
	// Functions that take a value argument must take a value of that specific type:
	v3 := Vertex{X: 3, Y: 4}
	p3 := &v3

	// var v Vertex
	// ScaleWithValueFunction(v, 5)  -> OK
	// ScaleWithValueFunction(&v, 5) -> Compile error!
	ScaleWithValueFunction(v3, 3)
	fmt.Println("Function call with pointer (v3):", v3)
	ScaleWithValueFunction(*p3, 3)
	fmt.Println("Function call with pointer (p3):", v3)

	// While methods with value receivers take either a value or a pointer as the receiver when they are called:
	// var v Vertex
	// v.ScaleWithValue() -> OK
	// p := &v
	// p.ScaleWithValue() -> OK
	v3.ScaleWithValue(3)
	fmt.Println("Value receiver method call (v3):", v3)
	p3.ScaleWithValue(3) // (*p3).ScaleWithValue(3)
	fmt.Println("Value receiver method call (p3):", v3)
	// In this case, the method call p.ScaleWithValue() is interpreted as (*p).ScaleWithValue()
}

// There are two reasons to use a pointer receiver:
// 1. The method can modify the value that its receiver points to.
// 2. To avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why, shortly).
// But if that's the case, can you tell what I did wrong above?
