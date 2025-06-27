package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eDo"))
	fmt.Println(regex.MatchString("e1o"))

	fmt.Println(regex.FindAllString("eko edo eDo e1o", 10))
}
