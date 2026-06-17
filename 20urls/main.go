package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=cgfgf6ffnb7"

func main() {
	fmt.Println("Welcome to urls")
	fmt.Println("url: ", myurl)
	fmt.Println("\n")
	result, _ := url.Parse(myurl)
	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query() //map used
	fmt.Println(qparams["paymentid"])

	for i, val := range qparams {
		fmt.Println(i, val)
	}

	partsOdUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherUrl := partsOdUrl.String()
	fmt.Println("Another url: ", anotherUrl)

}
 