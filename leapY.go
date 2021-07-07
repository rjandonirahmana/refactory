package main

import (
	"fmt"
)

func main() {

	hasil := leapYear(100, 200) // using struct and method, change its struct value by using method
	fmt.Println(hasil)

}

type year struct {
	key   int
	value []int
}

func (y *year) addYear(num int) {
	y.key = num
	y.value = append(y.value, y.key)

}

func leapYear(num, num1 int) []int {
	nilai := 0
	if num < num1 {
		nilai = num1 - num
	} else {
		nilai = num - num1
	}

	l := year{}

	for i := 0; i <= nilai; i++ {
		if i%4 == 0 {
			num2 := num + i
			l.addYear(num2)
		}
	}

	return l.value
}
