package main

import "fmt"

func main() {
	s := make([]int, 3)

	s[0] = 1
	s[1] = 2
	s[2] = 3

	s = append(s, 4)
	s = append(s, 'h')
	fmt.Println("emp:", s)

	s = append(s, []int{5, 6, 7}...)
	c := make([]int, len(s))
	fmt.Println("cpy:", len(c), "cap：", cap(c))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("l:", l)
	l = s[:5]
	fmt.Println("l:", l)
	l = s[2:]
	fmt.Println("l:", l)
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	towD := make([][]int, 3)
	fmt.Println("towD:", towD)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		towD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			towD[i][j] = i + j
		}
	}

	fmt.Println("towD:", towD)
	// for i := 0; i<3; i++{}
	// fmt.Println("set:", s)
	// fmt.Println("get:", s[2])

	// fmt.Println("len:", len(s))
	// s = append(s, 4)
	// fmt.Println("apd:", s)

}
