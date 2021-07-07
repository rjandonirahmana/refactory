package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Reverse("hallo Selamat datang"))
}

func Reverse(s string) string {
	kata := strings.Split(s, " ")

	rev := []string{}
	for _, v := range kata {
		r := []byte{}
		for i := 0; i < len(v); i++ {
			r = append(r, v[len(v)-i-1])
		}
		rev = append(rev, string(r))
	}

	hasil := strings.Join(rev, " ")
	return hasil

}
