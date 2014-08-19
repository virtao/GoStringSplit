package GoStringSplit

import (
	"container/list"
	"strings"
)

func SplitLine(line string) (result *list.List) {
	result = new(list.List)

	length := len(line)
	if length == 0 {
		return result
	}

	start := 0
	i := 0
	for ; i < length; i++ {
		switch line[i : i+1] {
		case " ":
			if start != i { //第一个字符为空格，或者连续2个或以上空格，会出现start和i相等的情况
				result.PushBack(line[start:i])
			}
			start = i + 1
		case "\"":
			if i+1 < length {
				pos := readUntil(line[i+1:], "\"")
				i += pos + 1
			}
		case "[":
			pos := readUntil(line[i:], "]")
			i += pos
		}
	}
	if start != i { //以空格结尾的字符串，会出现start等于i的情况。
		result.PushBack(line[start:i])
	}

	return result
}

func readUntil(s string, sep string) (pos int) {
	pos = strings.Index(s, sep)
	if pos == -1 {
		pos = len(s)
	}
	return pos
}
