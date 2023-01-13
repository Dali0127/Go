package main

import (
	"fmt"
	"strings"
)

func main() {
	var c1  complex64
	c1 = 1 + 2i
    fmt.Println(c1)

	s := "hello"
	s1 := "何力"

	// 字符串拼接
	fmt.Println(s + " " + s1)
	fmt.Printf("%s * %s \n", s , s1)

	// 字符串分割  strings.Split()
    s2 := "how old are you"
	fmt.Println(strings.Split(s2, " "))

	// %T 查看输出字符串的数据类型
	fmt.Printf("%T\n", strings.Split(s2, " "))

    // 判断是否包含
	fmt.Println(strings.Contains(s2, "do"))

	// 判断前缀
	fmt.Println(strings.HasPrefix(s2, "how"))

	// 判断后缀
	fmt.Println(strings.HasSuffix(s2, "you"))

	// 字符串出现的位置   (只取第一个值)
	fmt.Println(strings.Index(s2, "o"))
	// 字符串最后出现的位子
	fmt.Println(strings.LastIndex(s2, "o"))

	// join
	s3 := []string{"how", "do", "you", "do"}
	fmt.Println(s3)
    fmt.Println(strings.Join(s3, "-"))

	//byte

}
