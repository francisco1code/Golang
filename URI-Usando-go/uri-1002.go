package main

import (
	"fmt"
)
func main(){
	var pi float64 = 3.14159
	var a float64

	fmt.Scanf("%f",&a)

	var resultado float64

	resultado = a*a

	fmt.Printf("A=%.4f\n", resultado*pi)


}

