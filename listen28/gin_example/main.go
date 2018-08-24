package main

import "github.com/gin-gonic/gin"

func testHandle(c *gin.Context) {
	c.JSON(200,
		gin.H{"massage":"pong"})
}

func main(){
	//Default 返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//输出json调用方
		c.JSON(200,
			gin.H{"message":"pong"})
	})

	r.GET("/test", testHandle)
	r.Run(":9090")  //listen and server on 0.0.0.0:8080
}
