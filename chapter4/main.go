package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			fmt.Println("even")
		} else if i%3 == 0 {
			fmt.Println("divisible by 3")
		} else {
			fmt.Println("other number")
		}
	}
	for i := 1; i <= 3; i++ {
		switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		}
	}
}
