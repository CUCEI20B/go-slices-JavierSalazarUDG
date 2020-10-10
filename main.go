package main

import "fmt"

func main() {
	var s []int
	var length int
	fmt.Scan(&length)
	s = append(s, length)
	for i := 0; i < length; i++ {
		var value int
		fmt.Scan(&value)
		s = append(s, value)
	}
	sum := s[1 : length+1]
	var res int
	for _, v := range sum {
		res = res + v
	}
	print(res)
}
