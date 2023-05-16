package main

import "fmt"

func main() {
	fmt.Println("hello world")
	sprintf := fmt.Sprintf("%s,tmp.%d", "123", randomInt())
	fmt.Println(sprintf)
}

func randomInt() any {
	return 1
}
