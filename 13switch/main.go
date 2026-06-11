package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("Welcome to switch stuff")

	rand.Seed(time.Now().UnixNano())
	diceNumber:=rand.Intn(6)+1
	fmt.Println("Value of dice",diceNumber)


	switch diceNumber{
	case 1: 
	fmt.Println("Dice is 1 and you can open")
	case 2:
		fmt.Println("2 spot ahead")
		fallthrough
	case 3:
		fmt.Println("3 spot ahead")
		fallthrough
	case 4:
		fmt.Println("4 spot ahead")
		fallthrough
	case 5:
		fmt.Println("5 spot ahead")
		fallthrough
	case 6:
		fmt.Println("6 spot ahead and extra roll")
	default:
		fmt.Println("Errorrrrrrrrrrr")
	}
}