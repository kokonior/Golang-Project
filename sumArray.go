package main

import "fmt"

func main() {
	var arr = [...]int{2, 3, 2, 4, 3}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
