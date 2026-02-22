package main

import "fmt"

func typeOf() string {
	v := 42
	return fmt.Sprintf("%T",v)
}
