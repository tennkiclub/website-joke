// Package main provides ...
package main

import (
	"endpoints"
	"fmt"
	"github.com/gin-gonic/gin"
	eztemplate "github.com/michelloworld/ez-gin-template"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--runserver" {
			route := gin.Default()
			route.Static("/static", "./static")
			render := eztemplate.New()
			render.TemplatesDir = "template/"
			render.Ext = ".tmpl"
			render.Debug = true
			route.HTMLRender = render.Init()
			// route for Index
			route.GET("/", endpoints.Index)
			route.Run()
		} else {
			fmt.Println("Please use:")
			fmt.Println("--runserver : run gin Server. ")
			os.Exit(3)
		}
	} else {
		fmt.Println("Please use:")
		fmt.Println("--runserver : run gin Server. ")
		os.Exit(3)
	}
}
