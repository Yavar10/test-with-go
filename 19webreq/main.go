package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url="https://coderush-delta.vercel.app/"

func main()  {
	fmt.Println("Welcome to webreq")
	response, err :=http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response is of type %T",response)
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	defer response.Body.Close()
}