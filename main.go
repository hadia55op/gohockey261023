package main

import (
	"fmt"
	"hadia/gohockey261023/data"

	"github.com/gin-gonic/gin"
)

//	func hejhtmlhandlar(c *gin.Context) {
//		c.String(200, "<h1> hej!</h1><ul><li>Foppa</li><li> moppa</li></ul> ")
//	}
func hejhtmlhandlar(c *gin.Context) {
	c.Data(200, "text/html; charset=utf-8", []byte(`
        <h1>hej!</h1>
        <ul>
            <li>Foppa</li>
            <li>moppa</li>
        </ul>
    `))
}
func main() {
	// 	fmt.Println(" hello world")
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	r.GET("/hej.html", hejhtmlhandlar)
	if err := r.Run(":8081"); err != nil {
		fmt.Printf("failed to run server: %v", err)
	}
	fmt.Println(data.Hello())
}
