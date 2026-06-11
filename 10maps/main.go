package main

import "fmt"
func main()  {
	fmt.Println("Welcome to maps")
	languages:= make(map[string]string)

	languages["JS"]="JavaScript"
	languages["TS"]="TypeScript"
	languages["J"]="Java"
	languages["PY"]="Python"
	languages["RB"]="Ruby"

	fmt.Println("List of all languages: ",languages)
	fmt.Println("JS: ",languages["JS"])

	for _,value := range languages{
		fmt.Printf("For kay %v, value is %v\n",value)
	}

	for _,value := range languages{ //error ok syntax used here
		fmt.Printf("value is %v\n",value)
	}


}