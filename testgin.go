package main
import (
 "github.com/gin-gonic/gin"
 "strconv"
 // "net/http"
 // "fmt"
)
type User struct {
 Id int64 `db:"id" json:"id"`
 Firstname string `db:"firstname" json:"firstname"`
 Lastname string `db:"lastname" json:"lastname"`
}



// // func main() {
// //     router := gin.Default()

// //     // This handler will match /user/john but will not match neither /user/ or /user
// //     router.GET("/user/:name", func(c *gin.Context) {
// //         name := c.Param("name")
// //         c.String(http.StatusOK, "Hello %s", name)
// //     })

// //     // However, this one will match /user/john/ and also /user/john/send
// //     // If no other routers match /user/john, it will redirect to /user/john/
// //     router.GET("/user/:name/*action", func(c *gin.Context) {
// //         name := c.Param("name")
// //         action := c.Param("action")
// //         message := name + " is " + action
// //         c.String(http.StatusOK, message)
// //     })
// //     router.POST("/post", func(c *gin.Context) {

// //         id := c.Query("id")
// //         page := c.DefaultQuery("page", "0")
// //         name := c.PostForm("name")
// //         message := c.PostForm("message")

// //         fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
// //     })


// //     v1 := router.Group("/v1")
// //     {
// //         v1.POST("/login", loginEndpoint)
// //         v1.POST("/submit", submitEndpoint)
// //         v1.POST("/read", readEndpoint)
// //     }

// //     // Simple group: v2
// //     v2 := router.Group("/v2")
// //     {
// //         v2.POST("/login", loginEndpoint)
// //         v2.POST("/submit", submitEndpoint)
// //         v2.POST("/read", readEndpoint)
// //     }

// //     router.Run(":8080")
// // }


// // package main

// // import (
// //     "github.com/gin-gonic/gin"
// //     // "github.com/gin-gonic/gin/binding"
// // )

// // type LoginForm struct {
// //     User     string `form:"user" binding:"required"`
// //     Password string `form:"password" binding:"required"`
// // }

// // func main() {
// //     router := gin.Default()
// //     router.POST("/login", func(c *gin.Context) {
// //         // you can bind multipart form with explicit binding declaration:
// //         // c.BindWith(&form, binding.Form)
// //         // or you can simply use autobinding with Bind method:
// //         var form LoginForm
// //         // in this case proper binding will be automatically selected
// //         if c.Bind(&form) == nil {
// //             if form.User == "user" && form.Password == "password" {
// //                 c.JSON(200, gin.H{"status": "you are logged in"})
// //             } else {
// //                 c.JSON(401, gin.H{"status": "unauthorized"})
// //             }
// //         }
// //     })
// //     router.Run(":8080")
// // }

// func main() {
//     r := gin.Default()

//     // gin.H is a shortcut for map[string]interface{}
//     r.GET("/someJSON", func(c *gin.Context) {
//         c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
//     })

//     r.GET("/moreJSON", func(c *gin.Context) {
//         // You also can use a struct
//         var msg struct {
//             Name    string `json:"user"`
//             Message string
//             Number  int
//         }
//         msg.Name = "Lena"
//         msg.Message = "hey"
//         msg.Number = 123
//         // Note that msg.Name becomes "user" in the JSON
//         // Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
//         c.JSON(http.StatusOK, msg)
//     })

//     r.GET("/someXML", func(c *gin.Context) {
//         c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
//     })

//     // Listen and server on 0.0.0.0:8080
//     r.Run(":8080")
// }


func main() {
 r := gin.Default()
v1 := r.Group("api/v1")
 {
 v1.GET("/users", GetUsers)
 v1.GET("/users/:id", GetUser)
 v1.POST("/users", PostUser)
 v1.PUT("/users/:id", UpdateUser)
 v1.DELETE("/users/:id", DeleteUser)
 }
r.Run(":8080")
}


func GetUsers(c *gin.Context) {
  type Users []User
  var users = Users{
    User{Id: 1, Firstname: "Oliver", Lastname: "Queen"},
    User{Id: 2, Firstname: "Malcom", Lastname: "Merlyn"},
  }
  c.JSON(200, users)
  // curl -i http://localhost:8080/api/v1/users
}
func GetUser(c *gin.Context) {
 id := c.Params.ByName("id")
 user_id, _ := strconv.ParseInt(id, 0, 64)
if user_id == 1 {
 content := gin.H{"id": user_id, "firstname": "Oliver", "lastname": "Queen"}
 c.JSON(200, content)
 } else if user_id == 2 {
 content := gin.H{"id": user_id, "firstname": "Malcom", "lastname": "Merlyn"}
 c.JSON(200, content)
 } else {
 content := gin.H{"error": "user with id#" + id + " not found"}
 c.JSON(404, content)
 }
// curl -i http://localhost:8080/api/v1/users/1
}
func PostUser(c *gin.Context) {
 // The futur code…
}
func UpdateUser(c *gin.Context) {
 // The futur code…
}
func DeleteUser(c *gin.Context) {
 // The futur code…
}