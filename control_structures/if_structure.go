package control_structures

import "fmt"

func fizzBuzz(i int32) string {
	if i%3 == 0 && i%5 == 0 {
		return "FizzBuzz"
	}

	if i%3 == 0 {
		return "Fizz"
	}

	if i%5 == 0 {
		return "Buzz"
	}

	return fmt.Sprintf("%d", i)
}
