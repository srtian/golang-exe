package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goruntime")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	f("srtian")

	time.Sleep(time.Second)
	fmt.Println("done")
}