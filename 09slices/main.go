package main

import (
	"fmt"
	"sort"
)
func main(){
	fmt.Println("Welcome to the sliced stuff")
	var fruitlist = []string{"Apple","Mango","Banana"}
	//var fruitlist [8]string; array dec
	fmt.Printf("Type of fruitlist is %T \n",fruitlist)
	fruitlist=append(fruitlist, "Peach","Tomato")
	fruitlist= append(fruitlist[1:])
	fmt.Println(fruitlist)
	fruitlist= append(fruitlist[1:3])
	fmt.Println(fruitlist)
	
	fmt.Println("===================================================== ")

	highscore := make([]int, 4)  
	highscore[0]=245
	highscore[1]=55
	highscore[2]=2875
	highscore[3]=277
	//highscore[4]=628
	highscore=append(highscore,845,77,2,3,44)  //append allows to increase memory
	fmt.Println(highscore)


	fmt.Println("=====================================================")
	//sorting function
	fmt.Println("Is the slice sorted: ",sort.IntsAreSorted((highscore)))
	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println("Is the slice sorted: ",sort.IntsAreSorted((highscore)))
	
	fmt.Println("=====================================================")

	var courses =[]string{"reactjs","js","swift","python","ruby"}
	fmt.Println(courses)
	var index int =2
	courses= append(courses[:index],courses[index+1:]...)//skips val at index 
	// we needed to add ... cause appends signature is slice , elements so the slice had to be spread to elements to fit in the signature
	fmt.Println(courses)
}