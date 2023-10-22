package main

import (
	"fmt"
	"log"
)

// func main() {
// 	// nilPounter()
// 	// outOfRange()
// 	panicFunction()
// }

func panicFunction() {
	panic("error has occurred")
}

func nilPointer() {
	x := 10
	x_pointer := &x
	x_pointer = nil
	fmt.Print(*x_pointer)
}

func outOfRange() {
	names := []string{
		"lobster",
		"sea urchin",
		"sea cucumber",
	}
	fmt.Println("My favorite sea creature is:", names[len(names)])
}

// func main() {
// 	divideByZero()
// 	fmt.Println("we survived dividing by zero!")

// }

// func divideByZero() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.Println("panic occurred:", err)
// 		}
// 	}()
// 	fmt.Println(divide(1, 0))
// }

// func divide(a, b int) int {
// 	return a / b
// }

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("делитель равен 0")
	}
	return a / b, nil
}

func getPanic() {
	log.Panic("U MENYA PANIKA")
}

func main() {
	d, err := div(10, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("d = %d\n", d)
		getPanic()
	}
}
