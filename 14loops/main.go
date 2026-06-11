package main

import "fmt"

func main()  {
	fmt.Println("Welxome to loops")
	days:=[]string{"Sunday","Tuesday","Wed","Thurs","Fri","Sat"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i:=range days{
		fmt.Println(days[i])
	}
}