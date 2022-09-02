package main

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {
// 	content, err := ioutil.ReadFile("nikki225_module.json")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(string(content))
// }

// import (
// 	"fmt"
// 	j "test/jsonfile"
// )

// func main() {
// 	resutl := j.Readjsonfile("nikki225_module.json")
// 	fmt.Println(resutl)
// }

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Nikki225 struct {
// 	Id              int     `json:"id"`
// 	Code            int     `json:"code"`
// 	Long_name       string  `json:"long_name"`
// 	Short_name      string  `json:"short_name"`
// 	Weight          string  `json:"weight"`
// 	Last_price      float64 `json:"last_price"`
// 	Change_v        int     `json:"change_v"`
// 	Change_r        int     `json:"change_r"`
// 	Change_r_float  int     `json:"change_r_float"`
// 	Code_html       int     `json:"code_html"`
// 	Short_name_html int     `json:"short_name_html"`
// 	Type_html       int     `json:"type_html"`
// 	Market_value    int     `json:"market_value"`
// 	Return_R        int     `json:"return_R"`
// 	Year_max        int     `json:"year_max"`
// 	Year_min        int     `json:"year_min"`
// }

// func main() {
// 	msg := "{\"status\":200,\"msg\":18}"
// 	var data Nikki225
// 	if err := json.Unmarshal([]byte(msg), &data); err == nil {
// 		fmt.Println(data)
// 	} else {
// 		fmt.Println(err)
// 	}
// }

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/gin-gonic/gin"
	// j "test/jsonfile"
)

type Nikki225 struct {
	Id              int     `json:"id"`
	Code            int     `json:"code"`
	Long_name       string  `json:"long_name"`
	Short_name      string  `json:"short_name"`
	Type            string  `json:"type"`
	Weight          float64 `json:"weight"`
	Last_price      string  `json:"last_price"`
	Change_v        float64 `json:"change_v"`
	Change_r        string  `json:"change_r"`
	Change_r_float  float64 `json:"change_r_float"`
	Code_html       string  `json:"code_html"`
	Short_name_html string  `json:"short_name_html"`
	Type_html       string  `json:"type_html"`
	Market_value    string  `json:"market_value"`
	Return_R        string  `json:"return_R"`
	Year_max        string  `json:"year_max"`
	Year_min        string  `json:"year_min"`
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
	result := readjsonfile("nikki225_module.json", "device")
	fmt.Println(result)

	router := gin.Default()
	router.Static("/static", "templates/static")
	router.LoadHTMLFiles("templates/index.html")

	router.GET("/nikki225_module", func(c *gin.Context) {
		c.PureJSON(200, result)

	})

	router.StaticFile("index", "templates/index.html")
	router.StaticFile("img", "templates/nikki225_yearly.svg")
	router.Run()

}
