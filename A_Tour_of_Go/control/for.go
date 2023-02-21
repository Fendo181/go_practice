package main

import "fmt"

func main() {
	var sum,i int
	sum = 0
	for i = 0; i < 10; i++ {
		sum += i
	}
	sum = sum + 10
	fmt.Println(sum)
}
