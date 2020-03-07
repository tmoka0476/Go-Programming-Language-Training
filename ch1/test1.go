package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := os.Args[1]
	i, err := strconv.Atoi(n)
	fmt.Printf("%d,%x,%o,%b\n",i,i,i,i)
	fmt.Println(err)

}

