#strings — 字符串操作
```$xslt
字符串长度；
求子串；
是否存在某个字符或子串；
子串出现的次数（字符串匹配）；
字符串分割（切分）为[]string；
字符串是否有某个前缀或后缀；
字符或子串在字符串中首次出现的位置或最后一次出现的位置；
通过某个字符串将[]string连接起来；
字符串重复几次；
字符串中子串替换；
大小写转换；
Trim操作；
```
####1.1. 2.1.1 是否存在某个字符或子串
```$xslt
fmt.Println(strings.ContainsAny("team", "i"))
fmt.Println(strings.ContainsAny("failure", "u & i"))
fmt.Println(strings.ContainsAny("in failure", "s g"))
fmt.Println(strings.ContainsAny("foo", ""))
fmt.Println(strings.ContainsAny("", ""))
```
####1.2. 2.1.2 子串出现次数(字符串匹配)
```$xslt
fmt.Println(strings.Count("five", "")) // before & after each rune
```
####1.3.1. 2.1.3.1 Fields 和 FieldsFunc
```$xslt
fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
```
####1.3.2. 2.1.3.2 Split 和 SplitAfter、 SplitN 和 SplitAfterN
```$xslt
func Split(s, sep string) []string { return genSplit(s, sep, 0, -1) }
func SplitAfter(s, sep string) []string { return genSplit(s, sep, len(sep), -1) }
func SplitN(s, sep string, n int) []string { return genSplit(s, sep, 0, n) }
func SplitAfterN(s, sep string, n int) []string { return genSplit(s, sep, len(sep), n) }


fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))
```
####1.4. 2.1.4 字符串是否有某个前缀或后缀
```$xslt
// s 中是否以 prefix 开始
func HasPrefix(s, prefix string) bool {
    return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}
// s 中是否以 suffix 结尾
func HasSuffix(s, suffix string) bool {
    return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
```
####1.5. 2.1.5 字符或子串在字符串中出现的位置
```$xslt
// 在 s 中查找 sep 的第一次出现，返回第一次出现的索引
func Index(s, sep string) int
// chars中任何一个Unicode代码点在s中首次出现的位置
func IndexAny(s, chars string) int
// 查找字符 c 在 s 中第一次出现的位置，其中 c 满足 f(c) 返回 true
func IndexFunc(s string, f func(rune) bool) int
// Unicode 代码点 r 在 s 中第一次出现的位置
func IndexRune(s string, r rune) int

// 有三个对应的查找最后一次出现的位置
func LastIndex(s, sep string) int
func LastIndexAny(s, chars string) int
func LastIndexFunc(s string, f func(rune) bool) int
```
####1.6. 2.1.6 字符串 JOIN 操作
```$xslt
func Join(str []string, sep string) string {
    // 特殊情况应该做处理
    if len(str) == 0 {
        return ""
    }
    if len(str) == 1 {
        return str[0]
    }
    buffer := bytes.NewBufferString(str[0])
    for _, s := range str[1:] {
        buffer.WriteString(sep)
        buffer.WriteString(s)
    }
    return buffer.String()
}
```

```$xslt
func Join(a []string, sep string) string {
    if len(a) == 0 {
        return ""
    }
    if len(a) == 1 {
        return a[0]
    }
    n := len(sep) * (len(a) - 1)
    for i := 0; i < len(a); i++ {
        n += len(a[i])
    }

    b := make([]byte, n)
    bp := copy(b, a[0])
    for _, s := range a[1:] {
        bp += copy(b[bp:], sep)
        bp += copy(b[bp:], s)
    }
    return string(b)
}
```
####1.7. 2.1.7 字符串重复几次
```$xslt
fmt.Println("ba" + strings.Repeat("na", 2))
```
####1.8. 2.1.8 字符串子串替换
```$xslt
fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

// 输出
oinky oinky oink
moo moo moo
```
####1.9. 2.1.9 Replacer 类型