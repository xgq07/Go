package main

import "fmt"

func main() {
	a1 := [...]int{1, 2, 3}
	m1 := map[interface{}]string{
		a1: "windows",
	}
	fmt.Println(a1, m1)
	fmt.Println(m1[a1])
	fmt.Println(m1[[...]int{1, 2, 3}])

	var p1 *int
	v1 := 100
	p1 = &v1
	fmt.Println(v1, p1)
	m2 := map[interface{}]int{
		p1: 200,
	}
	fmt.Println(m2)
	fmt.Println(m2[p1])

	a2 := map[int]string{}
	fmt.Println(a2)
	fmt.Println(a2[100]) //不会panic

}
