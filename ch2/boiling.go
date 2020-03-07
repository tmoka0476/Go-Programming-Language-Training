package main

import "fmt"

func main() {
	fmt.Printf("boiling point = %gF or %gC\n",32.0,fToC(32.0))
}

func fToC (f float64) float64 {
	return (f-32)*5/9
}