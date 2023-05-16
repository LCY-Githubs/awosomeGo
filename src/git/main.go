package main

import "fmt"

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("there have a panic")
			fmt.Println(err)
		}
	}()
	fmt.Println("this is mini git")
	num1 := 10
	num2 := 0
	num3 := num1 / num2

	fmt.Println("num3=", num3)
}
