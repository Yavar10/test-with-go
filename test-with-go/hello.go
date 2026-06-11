package main
import "fmt"

func main2(){
	fmt.Println(Hello())
}
func Hello() string{
	return("Hello Zuri")
}
func Hello2(str string) string{
	return("Hello "+str)
}