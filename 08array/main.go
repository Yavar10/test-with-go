package main

import "fmt"

func main(){
	fmt.Println("Welcoome to array stuff")
	var fruitlist [40]string
	fruitlist[0]="aPPLE"
	fruitlist[1]="Tomato"
	fruitlist[3]="Peach"
	fmt.Println("Fruits: ",fruitlist) // leaves behind extra space between [1] and [3]
	fmt.Println("Length of the array: ",len(fruitlist));
}