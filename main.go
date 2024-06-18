package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func serverfile(g *gin.Context) {
	file := g.Param("filepath")
	dir := "view"
	if file == "/" {
		file = "/index.html"
	}
	files := strings.Split(file, ".")
	fmt.Println(files)
	fmt.Println(len(files))
	if len(files) <= 1 {
		file = file + ".html"
	} else if files[1] == "htm" || files[1] == "ht" {
		file = files[0] + ".html"
	} else if files[1] == "js" {
		file = file
	}
	fmt.Println(file)
	path := dir + file
	http.ServeFile(g.Writer, g.Request, path)
}
func server() {
	r := gin.Default()
	r.GET("/*filepath", serverfile)
	r.Run()
}
func main() {
	server()
}
