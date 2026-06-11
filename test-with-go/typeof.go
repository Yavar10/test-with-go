package main

import "fmt"

func typeOf() string {
	v := 42
	return fmt.Sprintf("%T",v)
}
func notmain(){
	var username string = "yavar"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n",isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n",smallVal)


}