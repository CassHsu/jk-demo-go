package main

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    hello := "Hello~"
    fmt.Println("Server starting ...")

    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, hello)
    })

    //listen port 8080
    r.Run()
}
