package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"sort"

	"github.com/gin-gonic/gin"
)
// https://github.com/coryb/sorty
//  map[string]interface{}-->sort
type sorter struct {
	data  interface{}
	order KeyComps
}

// NewSorter create a new sorter struct which Sort() can be called on.
func NewSorter() *sorter {
	return &sorter{}
}

// ByKeys is used to provide a list of string to indicate which keys to sort by and in which
// order.  If the key name starts with "-" it will be sorted in Descending order otherwise
// it will be sorted in Ascending order.
func (s *sorter) ByKeys(order []string) *sorter {
	keyComps := make(KeyComps, 0)
	for _, key := range order {
		switch key[0] {
		case '-':
			keyComps = append(keyComps, KeyComp{key[1:], Descending})
		case '+':
			keyComps = append(keyComps, KeyComp{key[1:], Ascending})
		default:
			keyComps = append(keyComps, KeyComp{key, Ascending})
		}
	}
	return s.ByKeyComps(keyComps)
}

// KeyComp struct to provide custom compaitor functions
type KeyComp struct {
	Name string
	Comp func(interface{}, interface{}) CompareResult
}

type KeyComps []KeyComp

// ByKeyComps is used to provide a list of KeyComp to sort by key with an explicit comparitor.
func (s *sorter) ByKeyComps(keyComps KeyComps) *sorter {
	s.order = keyComps
	return s
}

// Sort will sort the data provided.  The data should be a slice of something.
func (s *sorter) Sort(data interface{}) {
	s.data = data
	sort.Sort(s)
}

// Len is required to implement sort.Interface
func (s *sorter) Len() int {
	return reflect.ValueOf(s.data).Len()
}

// Swap is required to implement sort.Interface
func (s *sorter) Swap(i, j int) {
	if i > j {
		i, j = j, i
	}
	arr := reflect.ValueOf(s.data)

	tmp := arr.Index(i).Interface()
	arr.Index(i).Set(arr.Index(j))
	arr.Index(j).Set(reflect.ValueOf(tmp))
}

// Less is required to implement sort.Interface
func (s *sorter) Less(i, j int) bool {
	arr := reflect.ValueOf(s.data)
	a := reflect.ValueOf(arr.Index(i).Interface())
	b := reflect.ValueOf(arr.Index(j).Interface())
	if a.Kind() != reflect.Map {
		iface := a.Interface()
		panic(fmt.Sprintf("[A] Kind: %s, Expected a map, but got a %T for %v", a.Kind(), iface, iface))
	}
	if b.Kind() != reflect.Map {
		iface := b.Interface()
		panic(fmt.Sprintf("[B] Kind: %s, Expected a map, but got a %T for %v", b.Kind(), iface, iface))
	}

	for i := 0; i < len(s.order); i += 1 {
		keyComp := s.order[i]
		af := a.MapIndex(reflect.ValueOf(keyComp.Name)).Interface()
		bf := b.MapIndex(reflect.ValueOf(keyComp.Name)).Interface()

		switch keyComp.Comp(af, bf) {
		case LESSER:
			return true
		case GREATER:
			return false
		}
	}
	return true
}

type CompareResult int8

const (
	LESSER CompareResult = -1 + iota
	EQUAL
	GREATER
)

func Ascending(a, b interface{}) CompareResult {
	switch Descending(a, b) {
	case LESSER:
		return GREATER
	case GREATER:
		return LESSER
	default:
		return EQUAL
	}
}

func Descending(a, b interface{}) CompareResult {
	if a == b {
		return EQUAL
	}
	switch a.(type) {
	case string:
		return lg(a.(string) > b.(string))
	case int:
		return lg(a.(int) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int))
	case int8:
		return lg(a.(int8) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int8))
	case int16:
		return lg(a.(int16) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int16))
	case int32:
		return lg(a.(int32) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int32))
	case int64:
		return lg(a.(int64) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(int64))
	case uint:
		return lg(a.(uint) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint))
	case uint8:
		return lg(a.(uint8) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint8))
	case uint16:
		return lg(a.(uint16) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint16))
	case uint32:
		return lg(a.(uint32) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint32))
	case uint64:
		return lg(a.(uint64) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(uint64))
	case float32:
		return lg(a.(float32) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(float32))
	case float64:
		return lg(a.(float64) > reflect.ValueOf(b).Convert(reflect.TypeOf(a)).Interface().(float64))
	default:
		panic(fmt.Sprintf("dont know how to compare: %T", a))
	}
}

func lg(b bool) CompareResult {
	if b {
		return LESSER
	}
	return GREATER

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

func main() {
	// sort maps first in ascending order on the "foo" key
	// then in decending order on the "bar" key
	s := NewSorter().ByKeys([]string{
		"-change_r_float",
		// "-bar",
	})
	result := readjsonfile("nikki225_module.json", "device")

	s.Sort(result)
	fmt.Println(result)
	router := gin.Default()

	router.GET("/nikki225_module", func(c *gin.Context) {
		c.PureJSON(200, result)

	})
	router.Run()
}
