package main

import "fmt"

func main () {

	fmt.Println("hello")

	arr :=[]int{1,2,3,4}

	for _,i :=range arr {
		fmt.Println("i:",i)
	}

}
