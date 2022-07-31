package main

import (
	"fmt"
	"regexp"
)

const text = `邮箱地址为abc@163.com
bcd@qq.com
123@22.com
1@1.a.org
`

func main() {
	// regex, err := regexp.Compile("*@*") // 新建正则
	// if nil != err {
	// 	panic(err)
	// }

	re := regexp.MustCompile("([a-zA-Z0-9]*)@([a-zA-Z0-9]*\\.)+[a-zA-Z]+") // 这个不会返回错误，内部进行了err的处理  用两个斜杠的原因：\. 为golang中的转义字符
	// re := regexp.MustCompile(`*@.*\.*.`)  // 可以用`` 会鸳鸯输出

	match := re.FindString(text) // 在string中查找,只会查找第一个
	// re.Find() // 查找，参数为slice，可以把字符内穿转为slice放入

	matchs := re.FindAllString(text, -1) // 第二个参数为查找多少个 -1代表匹配所有, 不会把子匹配提取出来

	matchs_sub := re.FindAllStringSubmatch(text, -1) // 提取子匹配
	fmt.Println(match)
	fmt.Println(matchs)
	fmt.Println(matchs_sub)
}
