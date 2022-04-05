package main

import (
	"fmt"
)

func FilterInt(s []int, f func(int) bool) []int {
	var r []int
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

func FilterString(s []string, f func(string) bool) []string {
	var r []string
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

func main() {
	evens := FilterInt([]int{1, 2, 3, 4, 5}, func(i int) bool { return i%2 == 0 })
	fmt.Printf("%v\n", evens)

	shortStrings := FilterString([]string{"ok", "notok", "maybe", "maybe not"}, func(s string) bool { return len(s) < 3 })
	fmt.Printf("%v\n", shortStrings)
}
