package main

import "fmt"
func main(){
	fmt.Println("Welcome to a class on pointers")
	var ptr *int
	fmt.Println("Value of pointer is ",ptr)
	myNumber := 23
	 ptrv :=myNumber
	 ptr =&myNumber
	//var ptrv =myNumber
	fmt.Println("Value of pointer is ",ptrv)
	fmt.Println("Value of pointer ptr is ",*ptr)
	fmt.Println("Address of pointer ptr  is ",ptr)
}