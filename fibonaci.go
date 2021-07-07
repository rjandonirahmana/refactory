package main

import "fmt"

func main() {

	fmt.Println(searchNearest([]uint64{21}))
}

func Fibonacci(n int) uint64 {
	if n <= 1 {
		return uint64(n)
	}

	var n2, n1 uint64 = 0, 1

	for i := 0; i < n; i++ {
		n2, n1 = n1, n1+n2
		fmt.Println(n2, n1, "\n")
	}

	return n2 + n1
}

func searchNearest(nilai []uint64) uint64 {
	var n uint64
	for _, v := range nilai {
		n += v
	}

	slice := []uint64{}
	for i := 0; i <= 17; i++ {
		slice = append(slice, Fibonacci(i))
	}

	fmt.Println(slice)

	var nilaiSebelum uint64
	var nilaiRange uint64
	var nilaiSesudah uint64
	for i, value := range slice {

		if slice[i] == n {
			nilaiRange = slice[i]
			nilaiSesudah = slice[i+1]
			nilaiSebelum = slice[i-1]
		} else if slice[i] < n && slice[i+1] > n {
			nilaiRange = value
			nilaiSesudah = slice[i+1]
			nilaiSebelum = slice[i-1]
		}
	}
	fmt.Println(nilaiSebelum, nilaiRange, nilaiSesudah)

	if n == nilaiRange {
		if nilaiRange-nilaiSebelum > nilaiSesudah-nilaiRange {
			return nilaiSesudah - nilaiRange
		} else if nilaiRange-nilaiSebelum < nilaiSesudah-nilaiRange {
			return nilaiRange - nilaiSebelum
		}
	} else if n != nilaiRange {
		if n-nilaiRange > nilaiSesudah-n {
			return nilaiSesudah - n
		} else if n-nilaiRange < nilaiSesudah-n {
			return n - nilaiRange
		}
	}

	return n

}
