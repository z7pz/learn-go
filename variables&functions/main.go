package main

import "fmt"

func test(name string) string {
	var message string = fmt.Sprintf("Hello, %v", name)
	return message
}

func main() {
	println(test("Ahmed"))
}
