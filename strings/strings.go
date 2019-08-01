package main

import (
	"fmt"
	"strings"
)

func main() {
	// 子串substr在s中，返回true
	fmt.Println(strings.ContainsAny("team", "s"))
	fmt.Println(strings.ContainsAny("failure", "i & aaa"))
	fmt.Println(strings.ContainsAny("in failure", "sg"))
	fmt.Println(strings.ContainsAny("foo", "f"))
	//子串出现次数(字符串匹配)
	/**
	朴素匹配算法
	KMP算法
	Rabin-Karp算法
	Boyer-Moore算法
	*/
	//Rabin-Karp算法
	fmt.Println(strings.Count("five", ""))
	//字符串分割为[]string
	/**
	Fields 和 FieldsFunc、Split
	SplitAfter、SplitN
	SplitAfterN。
	*/
	fmt.Println(strings.Fields("foo bar  baz   "))

	//["foo" "bar" "baz"]
	fmt.Println(strings.Split("foo,bar,baz", ","))
	//["foo," "bar," "baz"]
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))
	fmt.Println(strings.SplitN("foo,bar,baz", ",", 2))

	// 1.4. 2.1.4 字符串是否有某个前缀或后缀
	// 子slice subslice 在 b 中，返回 true
	b := "123"
	fmt.Println(strings.Contains(b, "1"))

}
