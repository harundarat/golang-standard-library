package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	stringInt := strconv.Itoa(1000)
	fmt.Println(stringInt)

}
