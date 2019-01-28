package main

import (
	"regexp"
	"fmt"
)

const text = `My email is 931242644@qq.com@abc.com
email is abc@def.com
email is 	kkk@qq.com
email is ddd@agf.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match := re.FindAllString(text,-1)
	match := re.FindAllStringSubmatch(text, -1)
	for _,m := range match {
		fmt.Println(m)
	}
	//fmt.Println(match)
}
