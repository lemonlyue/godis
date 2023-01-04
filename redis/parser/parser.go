package parser

import (
	"bufio"
	"bytes"
	"io"
)

func parse0(rawReader io.Reader) {
	reader := bufio.NewReader(rawReader)
	for {
		// 获取文本协议
		line, err := reader.ReadBytes('\n')
		if err != nil {

		}

		length := len(line)
		if length <= 2 || line[length-2] != '\r' {
			// empty line
			continue
		}

		line = bytes.TrimSuffix(line, []byte{'\r', '\n'})
		// 字符开头
		switch line[0] {
		case '+': // 单行字符串
		case '-': // 错误提示
		case ':': // 整数值
		case '$': // 多行字符串
		case '*': // 数组
		default:

		}
	}
}
