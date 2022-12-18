package test

import "github.com/gin-gonic/gin"

func TestHello(c *gin.Context) {
	name := c.Request.URL.Path[len("/hello/"):]
	c.Writer.WriteString("Hello " + name)
}
