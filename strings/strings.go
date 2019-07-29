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
}
