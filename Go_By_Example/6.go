package main

import "fmt"


func add(nums ...int) int {
	fmt.Println(nums)
	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println(total)
	return total
}
func main() {
	add(1, 2, 3, 5)
	add(1, 2)
}