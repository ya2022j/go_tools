package main

import (
	"bytes"
	"fmt"
	"strings"
)

func cona_string1(str1, str2 string) (result string) {
	result = fmt.Sprintf("%s %s ", str1, str2)
	return result
}

func cona_string2(str1, str2 string) (result string) {
	result = str1 + str2
	return result
}

func cona_string3(str1, str2 string) (result string) {
	var str []string = []string{str1, str2}
	result = strings.Join(str, "")

	return result
}

func cona_string4(str1, str2 string) (result string) {
	var bt bytes.Buffer
	bt.WriteString(str1)
	bt.WriteString(str2)

	result = bt.String()

	return result
}

func cona_string5(str1, str2 string) (result string) {
	var build strings.Builder
	build.WriteString(str1)
	build.WriteString(str2)
	result = build.String()
	return result
}

func main() {
	s := cona_string4("hello", "world")
	fmt.Println(s)
}
