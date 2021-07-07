package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(fizzBuzz(20))

	fiz := fizzBuzz(30)
	fmt.Printf("%v", fiz)

}

func fizzBuzz(nilai int) []string {
	var array []string
	for i := 1; i < nilai; i++ {
		if i%3 == 0 && i%5 == 0 {
			str := strconv.Itoa(i)
			str = "FizzBuzz"
			array = append(array, str)

		} else if i%5 == 0 {
			st := strconv.Itoa(i)
			st = "Buzz"
			array = append(array, st)

		} else if i%3 == 0 {
			str := strconv.Itoa(i)
			str = "Fizz"
			array = append(array, str)

		} else {
			str := strconv.Itoa(i)
			array = append(array, str)
		}

	}
	return array
}
