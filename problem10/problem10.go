package main

import (
	"fmt"
)

func greatest(a ...int)(int){
	max := 0
	for _,i := range a{
		if i > max{
			max = i
		}
	}
	return max
}

func main(){
	m := greatest(1,2,3,4,5,6,7,8,9,10)
	fmt.Printf("The largest value was: %d\n",m)

	m = greatest(100,32,354,94,225,36,7,893,409,30)
	fmt.Printf("The largest value was: %d\n",m)
}