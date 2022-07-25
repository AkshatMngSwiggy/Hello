package main

import (
	"fmt"

	"github.com/AkshatMngSwiggy/greeting"
)

func main() {
	str := greeting.Hello("Akshat")
	fmt.Println("String Val : ", str)

	ret := num.Add(1, 3)
	fmt.Println("Return val : ", ret)
}
