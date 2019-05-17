package utils

import "regexp"

func ExtractString(contents []byte, re *regexp.Regexp) string {
	// 查找出第一个最匹配的结果
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
