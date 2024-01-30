package main

import "fmt"

func main() {
	var a int
	var b int
	a = 1
	b = 1
	fmt.Println("a + b=", a+b)
	var c float64
	var d float64
	fmt.Println("1.1 + 1.1 =", c+d)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	var helloWorld = "Hello world"
	fmt.Println(helloWorld)

	newString := "string"
	fmt.Println(newString)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println("x2", output)
}
