package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	old := intSeq()
	fmt.Println(old())
	fmt.Println(old())
	fmt.Println(old())
	new := intSeq()
	fmt.Println(new())
	fmt.Println(new())
}