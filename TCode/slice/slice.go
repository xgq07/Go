package main

import "fmt"

func main() {
	// 示例1。
	a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("a1: %v (len: %d, cap: %d)\n",
		a1, len(a1), cap(a1)) //[1...7], 7, 7
	s9 := a1[1:4]
	//s9[0] = 1
	fmt.Printf("s9: %v (len: %d, cap: %d)\n",
		s9, len(s9), cap(s9)) //2,3,4  3  6
	for i := 1; i <= 5; i++ {
		s9 = append(s9, i)
		fmt.Printf("s9(%d): %v (len: %d, cap: %d)\n",
			i, s9, len(s9), cap(s9)) //1 [2341]	4 6	; 2 [23412]5 6; 3 [234123]6 12; 4 7 8; 5 8 8
	}
	//注意，s9通过append扩容后 改变 数组a1的值了
	fmt.Printf("a1: %v (len: %d, cap: %d)\n",
		a1, len(a1), cap(a1)) //[1234123] ,7,7
	fmt.Println()

}
