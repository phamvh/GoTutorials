package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// See here for intro https://go.dev/tour/generics/1

// Index function returns the index of x in slice s, or -1 if not found
// The declaration of a generic type T must be in square brackets [], after func and before the arg list.
// In addition, here we also use the keyword comparable, which enforces a constraint that the values of
// type T are comparable with each other via == or !=.
// Without the keyword comparable, you cannot compare values of type T.
// Here, T is called a type parameter
//
//		T has another constraint, which says T has to be comparable, and comparable is an interface itself. This is often
//		called the meta type, or type constraint, meaning it imposes constraints on type parameter T.
//	    Details: https://go.dev/blog/intro-generics
func Index[T comparable](s []T, x T) int {
	for indx, val := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if val == x {
			return indx
		}
	}
	return -1
}

// Min returns the min of two elements of a generic type T
// Here, we also use constraints.Ordered to enforce that values of type T can be ordered, meaning they can
// be compared using > or <
// See https://go.dev/blog/intro-generics

func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// In general, we can use type parameter for struct also, not only for function as above.
// Let's do a linked list

// List is a linked list of generic type T
type List[T any /*any here means there is no constraint on type T*/] struct {
	next *List[T]
	val  T
}

// Length is a more complicated way of defining parameter type/**
// Here, S is of type "interface{~[]E}", which is a slice of elements of type E,
// and E, in turn, can be anything "E interface{}"
func Length[S interface{ ~[]E }, E interface{}](s S) int {
	return len(s)
}

// Length1 The above is equivalent to the following also
// Note that interface can be thought of as something containing either) a set of methods, or) a set of types
// See https://go.dev/blog/intro-generics
func Length1[S ~[]E, E any](s S) int {
	return len(s)
}

func main() {
	// We can use any type, which is comparable, as T, like int, float32, string, etc.
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(Index[int](slice1, 2)) // 1 // we put [int] here to tell the compiler that we use int as type
	// function argument type inference
	fmt.Println(Index(slice1, 2)) // here we OMIT [int], because Go can infer the type  int from the args passed in.

	slice2 := []string{"hello", "hi", "bonjour"}
	fmt.Println(Index(slice2, "hola")) // -1

	fmt.Println(Min(2, 4))            // 2
	fmt.Println(Min("Hello", "Hola")) // Hello
}
