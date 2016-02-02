package main

import (
	"fmt"
)

func half(a int)(int,bool){
	return a/2,a%2 == 0
}

func main(){
	a,b := 10,11
	c,d := half(a)
	fmt.Printf("Half of %d is %d\n",a,c)
	if(d) {
		fmt.Printf("%d is even\n", a)
	}else{
		fmt.Printf("d is not even\n",a)
	}
	c,d = half(b)
	fmt.Printf("Half of %d is %d\n",b,c)
	if(d) {
		fmt.Printf("%d is even\n", b)
	}else{
		fmt.Printf("%d is not even\n",b)
	}
}
