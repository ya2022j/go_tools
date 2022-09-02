package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// golang中文件和路径用法
// https://www.cnblogs.com/mayanan/p/15416565.html
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

func readjsonfile(jsonfile string, pathfile string) (result interface{}) {

	devicePath, _ := filepath.Abs(pathfile)
	jsonFilePath := filepath.Join(devicePath, jsonfile)

	content, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		panic(err)
	}
	// var str_content string
	string_result := string(content)
	// for_gin_json := map[string]interface{}{"json_result": result}

	//string--->json array

	// var data Nikki225

	var data []map[string]interface{}
	if err := json.Unmarshal([]byte(string_result), &data); err == nil {
		return data
		//fmt.Println(dat["status"])
	} else {
		fmt.Println(err)
	}
	return
}

func concatenate_string(str1, str2 string) (result string) {
	var build strings.Builder
	build.WriteString(str1)
	build.WriteString(str2)
	result = build.String()
	return result
}

func main() {
	s := readjsonfile("nikki225_module.json", "device")
	fmt.Println(s)
}
