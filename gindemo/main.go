package main

import (
	"net/http"
    "fmt"
    "time"
    "log"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        //fmt.Println(c)
        c.JSON(200, gin.H{
            "data": "hello  " + name,
        })
    })

    r.GET("/info", func(c *gin.Context) {
        p1 := c.Query("p1")
        p2 := c.Query("p2")
        //fmt.Println(c)
        data := fmt.Sprintf("p1:%s p2:%s", p1, p2)
        c.JSON(http.StatusOK, gin.H{
            "data": data,
        })
    })

    // 异步处理
    r.GET("/async", func(c *gin.Context) {
        copyCtx := c.Copy()
        go func() {
            time.Sleep(1 * time.Second)
            log.Println("异步执行：" + copyCtx.Request.URL.Path)
            // copyCtx.JSON(200, gin.H{
            //     "data": "async",
            // })
        }() 
    })

    // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
    r.Run()
    //r.Run(":8081") 
}