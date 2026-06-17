package main

import "fmt"
func main()  {
	fmt.Println("Welcome to methods")
//no inheritance no super no parent in golang
	yavar:=User{"Yavar","yavar@gmail.com",true,21}
	fmt.Println(yavar)
	fmt.Printf("Details: %v\n",yavar)
	fmt.Printf("Details: %+v\n",yavar)
	fmt.Println(yavar.Name," is old by ",yavar.Age)
	fmt.Println("\n")
	yavar.GetStatus()
	yavar.PrintMail()
	yavar.NewMail()//stll doesnt change the original caaause only copy was passed
	yavar.PrintMail()
}

type User struct {
	Name   string
	Email  string
	Status bool         //capital letter public thing 
	Age    int
	//oneAge int          //not exportable 
}


func (u User) GetStatus() {
	fmt.Println("Is user active: ",u.Status)
}

func (u User) NewMail() {
	u.Email="test@go.dev"
}
func (u User) PrintMail() {
	fmt.Println("\n",u.Email)
}


