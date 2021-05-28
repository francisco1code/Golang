package main

import (
	"fmt"
)

func main(){

	var a int = 0
	var b int = 0

	fmt.Scan(&a, &b)

	var soma int = a+b
	
	fmt.Printf("X = %d\n",soma)
}