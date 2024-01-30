package main

import "fmt"

func calculateTotal(arr []float64) float64 {
	var total float64 = 0
	for _, value := range arr {
		total += value
	}
	return total
}

func main() {
	arr := []float64{98, 91, 55, 65, 100}
	var total = calculateTotal(arr)
	fmt.Println(total)
	var sliceTotal = calculateTotal(arr[2:4])
	fmt.Println(sliceTotal)
}
