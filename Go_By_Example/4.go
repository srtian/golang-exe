package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "a"
    s[1] = "b"
    s[2] = "c"
	fmt.Println(s)
	s = append(s, "d")
    s = append(s, "e", "f")
	fmt.Println("apd:", s)
	l := s[1:2]
	fmt.Println("l", l)
}