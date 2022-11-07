package main

import "github.com/gin-gonic/gin"

func testHello(c *gin.Context) {
	name := c.Request.URL.Path[len("/hello/"):]
	c.Writer.WriteString("Hello " + name)
}

func lineCallback(c *gin.Context) {

}
