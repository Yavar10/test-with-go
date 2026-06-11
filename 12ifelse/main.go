package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main()  {
	fmt.Println("Welcome to if-else")
	fmt.Println("Enter a number")
	//reader := bufio.NewReader(os.Stdin) 
	reader :=bufio.NewReader(os.Stdin)
	input,_:=reader.ReadString('\n')

	loginCount,_:=strconv.ParseInt(strings.TrimSpace(input),10,64)
	var result string
	
	if loginCount < 10 {
		result="Regular user"
	} else if loginCount==10{ 
        result="watch Out"
	} else{
		result="overflow"
	}

	fmt.Println(result)

}