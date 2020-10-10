package main

import "fmt"

func main() {
	var s []float32
	var length int
	fmt.Scan(&length)
	s = append(s, float32(length))
	for i := 0; i < length; i++ {
		var value float32
		fmt.Scan(&value)
		s = append(s, value)
	}
	sum := s[1 : length+1]
	var res float32
	for _, v := range sum {
		res = res + v
	}
	fmt.Printf("%f\n", res)
}
