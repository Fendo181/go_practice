package main

import (
	"fmt"
	"math"
)

func pow(x,n,lim float64) float64 {
	// べき乗の値が(xのn乗)がlimよりも小さい場合はべき乗の値を返す
	if v := math.Pow(x,n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}
