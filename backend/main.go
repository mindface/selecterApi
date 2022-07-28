package main

import (
 // "net/http"
 "fmt"
 "github.com/app/controller"
 "github.com/gin-contrib/cors"
//  _ "github.com/go-sql-driver/mysql"
//  _ "github.com/jinzhu/gorm"
 "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello")
	r := gin.Default()

	r.Use(cors.New(cors.Config{
  AllowOrigins:     []string{"http://localhost:3001"},
  AllowMethods:     []string{"POST", "PUT", "PATCH"},
  AllowHeaders:     []string{"Content-Type"},
  AllowCredentials: true,
}))

r.GET("/api",func(c *gin.Context){
 c.JSON(200, gin.H {"message":"hello"})
 })

r.GET("/api2",func(c *gin.Context){
 c.JSON(200, gin.H {"message":"hello22"})
 })

 r.GET("/ddd",func(c *gin.Context){
  nllk := controller.Sss()
  c.JSON(200, gin.H{"message":nllk})
  })
 // r.GET("/cardsinit", controller.IndexInit)
 // r.GET("/cards", controller.IndexGet)
 // r.POST("/cards", controller.IndexPost)

r.POST("/api_post", func(c* gin.Context) {
  msg := c.PostForm("msg")
  nick := c.DefaultPostForm("nick","anoymous")
  fmt.Println("------------")
  fmt.Println(msg)

  c.JSON(200,gin.H{
    "status":"posted",
    "msg": msg,
    "nick": nick,
  })
})

 r.Run(":3000")

}