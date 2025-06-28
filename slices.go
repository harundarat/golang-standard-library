package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 96, 80, 90}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Contains(names, "Harun"))
	fmt.Println(slices.Index(names, "Harun"))
	fmt.Println(slices.Index(names, "Paul"))

}
