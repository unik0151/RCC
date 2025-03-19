package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/unik0105/go-demo1/db"
	_ "github.com/unik0105/go-demo1/db"
	"github.com/unik0105/go-demo1/utools"
	_ "github.com/unik0105/go-demo1/utools"

	"github.com/gin-gonic/gin"
	_ "github.com/unik0105/go-demo1/db"
)

// 如果多个html只会加载一个
// 构建常用的传参:

func main() {
	router := gin.Default()
	// 多个加载
	// 1. loadMultiHtml(router)
	// 2. 页面加入函数处理
	// addConvertFunc02(router)
	// 3.上传
	uploadFile(router)
	// 4. 将图片上传之后并展示
	selectFileToPage(router)
	router.Run(":8080")
}

func loadMultiHtml01(router *gin.Engine) {

	router.LoadHTMLGlob("templates/**/*")
	router.SetTrustedProxies(nil)
	// 多个func会追加
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	}, func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "posts/post.tmpl", gin.H{
			"post": "post",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
}

// 为页面内置转换函数
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()

	fmt.Println("year", year, " month:", month, " day: ", day)
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func addConvertFunc02(router *gin.Engine) {
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLGlob("templates/row/*")
	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "row.html", map[string]any{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

}

var logName string = "go.log"

var relativePath = "file/"

func uploadFile(router *gin.Engine) {
	file, err := os.Create(logName)
	if err != nil {
		panic(err.Error())
	}

	gin.DefaultWriter = io.MultiWriter(file)
	router.MaxMultipartMemory = 8 << 20

	fileV1 := router.Group("/file")
	{
		fileV1.Handle(http.MethodPost, "/add", func(ctx *gin.Context) {
			var typeA string = ctx.GetHeader("Content-Type")
			log.Println("typeA: ", typeA)
			formFile, err := ctx.MultipartForm()
			files := formFile.File["upload[]"]
			if err != nil {
				panic(err.Error())
			}
			for _, file := range files {
				src, _ := file.Open()
				defer src.Close()
				fullName := relativePath + file.Filename
				utools.CreateIfExistNo(fullName)
				fileVar := db.FileDTO{Creator: "chen", FilePath: fullName}
				result := db.CreateDATA(db.GetDBDriver(), &fileVar)
				fmt.Printf("result.RowsAffected: %v\n", result.RowsAffected)
				fmt.Printf("result.Error: %v\n", result.Error)
				dst, _ := os.Create(fullName)
				defer dst.Close()
				if _, err := io.Copy(dst, src); err != nil {
					panic("err :" + err.Error())
				}

			}
		})
	}
}

func selectFileToPage(router *gin.Engine) {

	v1 := router.Group("/img")
	{
		v1.GET("/get/:id", func(ctx *gin.Context) {
			var name = ctx.Param("id")
			ctx.Header("Content-Type", "image/png")
			name = relativePath + name
			if _, err := os.Stat(name); err != nil {
				ctx.String(400, "current file is not exists!")
			}
			bytes, err := os.ReadFile(name)
			if err != nil {
				panic("读取失败")
			}
			ctx.Writer.Write(bytes)
		})
	}
}

func createTimeStamp() int64 {
	return time.Now().Unix()
}
