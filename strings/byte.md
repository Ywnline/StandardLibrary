#bytes — byte slice 便利操作

###是否存在某个子slice
```$xslt
// 子slic subslice 在 b 中是否存在，返回true
func Contains(b, subslice []byte) bool
```

###[]byte 出现次数
```$xslt
// slice sep 在 s 中出现的次数（无重叠）
func Count(s, sep []byte) int
```