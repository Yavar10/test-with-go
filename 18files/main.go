package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
func main()  {
	fmt.Println("Welcome to files")
	content := "This needs to go in a file "

	file,err:=os.Create("./mylcogofile.txt")
/* 
	if err != nil {
		panic(err)
	} */
	checkNilErr(err)

	length , err :=io.WriteString(file ,content)
	checkNilErr(err)
	 fmt.Println(length)
	 defer file.Close()
	 readFile("./mylcogofile.txt")
}

func readFile(filname string)  {
	databyte, err :=ioutil.ReadFile(filname)
	checkNilErr(err)
	fmt.Println((databyte))
	fmt.Println(string(databyte))
}
func checkNilErr(err error)  {
	if err != nil {
		panic(err)
	}
}

