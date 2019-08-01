package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []int{2, 1, 4, 5, 3, 7, 9, 8, 6}
	//正序
	sort.Ints(s)
	// 倒序  IntSlice
	//sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
	// SearchInts 必须为升序数组
	fmt.Println(sort.SearchInts(s, 9)) //将会输出0而不是1

	/**
	2. Float64Slice类型及[]float64排序
	....
	....
	....
	3. StringSlice类型及[]string排序
	....
	....
	....
	*/

}
