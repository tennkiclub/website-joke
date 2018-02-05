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
			gin.SetMode(gin.ReleaseMode)
			route := gin.Default()
			route.Static("/static", "./static")
			render := eztemplate.New()
			render.TemplatesDir = "template/"
			render.Ext = ".tmpl"
			route.HTMLRender = render.Init()
			// route for Index
			route.GET("/", endpoints.Index)
			route.Run(":8509")
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
