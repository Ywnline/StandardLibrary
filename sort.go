package StandardLibrary

import (
	"fmt"
	"sort"
)

//学生成绩结构体
type StuScore struct {
	name  string // 学生
	score int    // 成绩
	sge   int    // 年龄
}
type StuScores []StuScore

//len()
func (s StuScores) Len() int {
	return len(s)
}

//Less():成绩将有低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].sge < s[j].sge
}

//Swap()
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"alan", 95, 11},
		{"hikerell", 91, 12},
		{"acmfly", 96, 22},
		{"leao", 90, 21},
	}
	//打印未排序的stus数据
	fmt.Println("Default:\n\t", stus)
	//StuScores已经实现了sort.Interface接口,所以可以调用Sort函数进行排序
	sort.Sort(stus)
	//判断是否已经排好顺序，将会打印true
	fmt.Println("IS Sorted?\n\t", sort.IsSorted(stus))
	//打印排序后的stus数据
	fmt.Println("Sorted:\n\t", stus)
}
