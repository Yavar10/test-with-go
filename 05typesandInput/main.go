package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	welcome:= "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) //Scanner sc = new Scanner(System.in)
	fmt.Println("Enter Pizza rating: ")
	//comma okk // error ok
	//we have no try catch
	input,_:=reader.ReadString('\n');
	//input,err:=reader.ReadString('\n');
	//_,err:=reader.ReadString('\n');
	fmt.Printf("Rating: %v\n",input)
	fmt.Printf("Type: %T\n",input)

	//conversion
	numRating,err:=strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to rating: ",numRating+1)
	}
	
}