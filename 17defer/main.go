package main

import "fmt"

func main()  {
	defer fmt.Println("world") //lifo 
	defer fmt.Println("world2") //lifo 
	fmt.Println("hello")
	MyDefer()
}

func MyDefer()  {
	defer fmt.Println("mD1") //diff lifo first this defer exe then outside
	defer fmt.Println("mD2") 
	fmt.Println("method")
	
}