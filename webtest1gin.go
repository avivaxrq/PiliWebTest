package main
import (
 "github.com/gin-gonic/gin"
 // "strconv"
 "net/http"
 // "fmt"
)
func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "webadmin.tmpl", gin.H{
            "title": "piliwebtest",
            ""
        })
    })
    // router.GET("/users/index", func(c *gin.Context) {
    //     c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
    //         "title": "Users",
    //     })
    // })
    router.Run(":8080")
}