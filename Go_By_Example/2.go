package main

import (
	"fmt"
	"time"
)

func main() {
	 i := 2
	 fmt.Print("write", i,  " as ")
	 switch i {
	 case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	 }

	 switch time.Now().Weekday() {
	 case time.Saturday, time.Sunday:
		fmt.Println("It is Weekend")
	default:
		fmt.Println("It is weekday")
	 }
	 t := time.Now()
	 switch {
	 case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	 }
	 whatAmI := func(i interface{}) {
		 switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("i don't know %T\n", t)
		 }
	 }
	 whatAmI(true)
	 whatAmI(1)
	 whatAmI("hh")
}