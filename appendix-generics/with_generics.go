package main

import (
	"fmt"
)

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

func main() {
	evens := Filter([]int{1, 2, 3, 4, 5}, func(i int) bool { return i%2 == 0 })
	fmt.Printf("%v\n", evens)

	shortStrings := Filter([]string{"ok", "notok", "maybe", "maybe not"}, func(s string) bool { return len(s) < 3 })
	fmt.Printf("%v\n", shortStrings)
}
