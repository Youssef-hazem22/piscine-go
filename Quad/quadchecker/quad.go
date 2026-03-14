package main

import (
	"fmt"
	"os"

	"quadchecker/quad"
)

func normalize(data []byte, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		ch := data[i]
		if ch == '\r' {
			continue
		}
		result += string(ch)
	}
	start := 0
	for start < len(result) && (result[start] == ' ' || result[start] == '\n') {
		start++
	}
	end := len(result) - 1
	for end >= 0 && (result[end] == ' ' || result[end] == '\n') {
		end--
	}
	if start > end {
		return ""
	}
	return result[start : end+1]
}

func main() {
	data := make([]byte, 10000)
	n, _ := os.Stdin.Read(data)
	shape := normalize(data, n)

	if len(shape) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	matches := []string{}

	for y := 1; y <= 50; y++ {
		for x := 1; x <= 50; x++ {
			if normalize([]byte(quad.QuadA(x, y)), len(quad.QuadA(x, y))) == shape {
				matches = append(matches, "[quadA] ["+fmt.Sprint(x)+"] ["+fmt.Sprint(y)+"]")
			}
			if normalize([]byte(quad.QuadB(x, y)), len(quad.QuadB(x, y))) == shape {
				matches = append(matches, "[quadB] ["+fmt.Sprint(x)+"] ["+fmt.Sprint(y)+"]")
			}
			if normalize([]byte(quad.QuadC(x, y)), len(quad.QuadC(x, y))) == shape {
				matches = append(matches, "[quadC] ["+fmt.Sprint(x)+"] ["+fmt.Sprint(y)+"]")
			}
			if normalize([]byte(quad.QuadD(x, y)), len(quad.QuadD(x, y))) == shape {
				matches = append(matches, "[quadD] ["+fmt.Sprint(x)+"] ["+fmt.Sprint(y)+"]")
			}
			if normalize([]byte(quad.QuadE(x, y)), len(quad.QuadE(x, y))) == shape {
				matches = append(matches, "[quadE] ["+fmt.Sprint(x)+"] ["+fmt.Sprint(y)+"]")
			}
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	for i := 0; i < len(matches); i++ {
		if i > 0 {
			fmt.Print(" || ")
		}
		fmt.Print(matches[i])
	}

}
