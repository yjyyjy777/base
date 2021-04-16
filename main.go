package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
func main(){
	r:=gin.Default()
	r.GET("hello/:id/:qq", func(c *gin.Context) {
		name:=c.DefaultQuery("name","hello")
		par :=c.Param("id")
		par2:=c.Param("qq")
		fmt.Println(par)
		c.Writer.WriteString(par+" "+par2+" "+name)
	})
	r.POST("login", func(c *gin.Context) {
		path :=c.FullPath()
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		fmt.Println(path)
		c.Writer.Write([]byte("hello"+" "+password+" "+ username))
	})


	r.Run()
}

