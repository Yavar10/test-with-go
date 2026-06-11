package main

import "fmt"

func main() {
	fmt.Println("Welcome tp structs")
	//no inheritance no super no parent in golang
	yavar:=User{"Yavar","yavar@gmail.com",true,21}
	fmt.Println(yavar)
	fmt.Printf("Details: %v\n",yavar)
	fmt.Printf("Details: %+v\n",yavar)
	fmt.Println(yavar.Name," is old by ",yavar.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool         //capital letter public thing 
	Age    int
}
