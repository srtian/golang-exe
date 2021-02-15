package main

import "fmt"

func val(v int) {
	v = 0
}

func ptr(v *int) {
	*v = 0
}

func main() {
	i := 1
	val(i)
	fmt.Println(i)
	ptr(&i)
	fmt.Println(i)
	fmt.Println("pointer:", &i)
}