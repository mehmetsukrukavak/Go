package main

import (
	"fmt"
	//"os"
	"strconv"
)

func main() {
	fmt.Println("Hello World!")
/*
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
*/
	var myString = "1"

	number, _ := strconv.Atoi(myString)

	fmt.Println("Sonuç : " +strconv.Itoa(number))



}
