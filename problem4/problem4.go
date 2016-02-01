package main

import (
	"fmt"
	"log"
)

func main(){

	var big, small, remainder int
	var err error
	fmt.Print("Enter a large number and a small number separated by a space: ")

	fmt.Scanf("%d%d",&big,&small,err)

	if(err != nil){
		log.Fatal(err)
	}

	if(big < small){
		remainder = big
		big = small
		small = remainder
	}

	remainder = big % small
	fmt.Printf("The remainder of %d divided by %d is %d \n",big,small,remainder)
}