package main

import (
	"fmt"
)

func main(){

	i := 1

	for i < 100 {
		if i % 2 == 0{
			fmt.Printf("%d ",i)
		}
		i++
	}
	fmt.Print("\n")
}