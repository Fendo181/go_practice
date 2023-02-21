package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		// 0より下回って場合は平方根にする
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
