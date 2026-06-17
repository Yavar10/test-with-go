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

	for index,day:=range days{
		fmt.Println(index," ",day)
	}

	for _,day:=range days{
		fmt.Println(day)
	}
	hday:=1
	for hday<10 {
		fmt.Println("yes")
		hday++
	}
}