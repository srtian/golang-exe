package main

import ("fmt")

func main() {
	a := make(map[string]int)
	a["id"] = 11
	a["ids"] = 111
	fmt.Println(a)
	delete(a, "ids")
	fmt.Println(a)
	_, is := a["id"]
	fmt.Println(is)
}