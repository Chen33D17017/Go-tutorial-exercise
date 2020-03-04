package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (m map[string]int) {
	m = make(map[string]int)
	test := strings.Split(s, " ")
	for _, sub_string:= range test{
		map_value, exist := m[sub_string]
		if exist {
			m[sub_string] = map_value+1
		} else {
			m[sub_string] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
