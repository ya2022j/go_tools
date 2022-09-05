package main

import (
	"net/http"
	"strings"
	"test/routes"

	"github.com/gin-gonic/gin"
)

//编译时出现了其他文件所以之前总是编译不过去
// 路由管理也就是一个路由对应一个视图函数
// 目前把html解决了，回头处理json。正常也就可以使用了

func cona_string5(str1, str2 string) (result string) {
	var build strings.Builder
	build.WriteString(str1)
	build.WriteString(str2)
	result = build.String()
	return result
}
func Jp(c *gin.Context) {
	id := c.Param("id")
	id_html := cona_string5(id, ".html")
	c.HTML(http.StatusOK, id_html, gin.H{
		"id": id,
	})

}

func NoRoute(c *gin.Context) {
	// helloに飛ばす
	c.Redirect(http.StatusMovedPermanently, "/jp")
}
func main() {
	//os.Setenv("GIN_MODE", "release")
	//gin.SetMode(gin.ReleaseMode)

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/jp/:id", Jp)

	router.NoRoute(routes.NoRoute) // どのルーティングにも当てはまらなかった場合に処理

	router.Run(":8080")
}
