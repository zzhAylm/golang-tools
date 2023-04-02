package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	//下标 1 到 3 ，最后一个不要
	a := s[1:4]
	fmt.Println(a)
	//下标 0 到 1 ，最后一个不要
	b := s[:2]
	fmt.Println(b)
	//下标 0 到 最后一个 ，最后一个不要
	c := s[0:]
	fmt.Println(c)
}
