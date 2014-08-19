package GoStringSplit

import (
	"container/list"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SplitLine(t *testing.T) {
	testStr := ""
	result := new(list.List)
	result = SplitLine(testStr)
	assert.NotNil(t, result, "测试解析空字符串，解析结果不应为nil。")
	ResultOutput(result)
	assert.Equal(t, 0, result.Len(), "测试解析空字符串，解析结果应该为0个元素。")

	testStr = `115.231.14.83 - - [18/Aug/2014:11:39:42 +0800] "GET /robots.txt HTTP/1.1" 0.231 200 24 "-" "ChinaCache" "-"`
	result = nil
	result = SplitLine(testStr)
	assert.NotNil(t, result, "测试解析正常字符串，解析结果不应为nil。")
	ResultOutput(result)
	assert.Equal(t, 11, result.Len(), "测试解析正常字符串，解析结果应该为11个元素。")

	testStr = `42.159.4.75 - - [18/Aug/2014:11:39:41 +0800] "GET /1000/797_612-411.jpg HTTP/1.1" 0.191 304 0 "http://www.xxx.com/backend/location?page=39" "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36" "123.207.88.67"`
	result = nil
	result = SplitLine(testStr)
	assert.NotNil(t, result, "测试解析正常字符串，解析结果不应为nil。")
	ResultOutput(result)
	assert.Equal(t, 11, result.Len(), "测试解析正常字符串，解析结果应该为11个元素。")

	testStr = ` 115.231.14.83 - - [18/Aug/2014:11:39:42 +0800] "GET /robots.txt HTTP/1.1" 0.231 200 24 "-"  "ChinaCache" "-" `
	result = nil
	result = SplitLine(testStr)
	assert.NotNil(t, result, "测试解析错误的字符串，字符串包含连续的空格以及首尾空格，解析结果不应为nil。")
	ResultOutput(result)
	assert.Equal(t, 11, result.Len(), "测试解析错误的字符串，字符串包含连续的空格以及首尾空格，解析结果应该为11个元素。")

	testStr = `11"5.231.14.83 - -" [18/Aug/2014:11:39:42 +0800] ""GET /robots.txt HTTP/1.1" 0.231 200 24 "-" "ChinaCache" "-""`
	result = nil
	result = SplitLine(testStr)
	assert.NotNil(t, result, "测试解析错误的字符串，字符串包含不正确的引号，解析结果不应为nil。")
	ResultOutput(result)
	assert.Equal(t, 5, result.Len(), "测试解析错误的字符串，字符串包含不正确的引号，解析结果应该为5个元素。")

}

func ResultOutput(result *list.List) {
	for e := result.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
