package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Harun Al Rasyid", "run"))
	fmt.Println(strings.Split("Harun Al Rasyid", " "))
	fmt.Println(strings.ToLower("Harun Al Rasyid"))
	fmt.Println(strings.ToUpper("Harun Al Rasyid"))
	fmt.Println(strings.Trim("   Harun Al Rasyid    ", " "))
	fmt.Println(strings.Replace("Harun Rasyid Al Rasyid", "Rasyid", "Irsyad", -1))
}
