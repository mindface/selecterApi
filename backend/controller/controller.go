package controller

import (
		_ "github.com/go-sql-driver/mysql"
	 // _ "github.com/jinzhu/gorm"
		"gorm.io/gorm"
)

type JsonRequest struct {
	 gorm.Model
		Title string `json:"title"`
		Description string `json:"description"`
}

type Card struct {
	 gorm.Model
		Title string `json:"title"`
	Description string `json:"description"`
}

// func gormConnect() *gorm.DB {
// 	DBMS := "mysql"
// 	USER := "gowebdb"
// 	PASS := "gowebdb"
// 	DBNAME := "gowebdb"
// 	CONECT := USER + ":" + PASS + "@tcp(127.0.0.1:3306)/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"

// 	db, err := gorm.Open(DBMS,CONECT)
// 	if err != nil {
// 		 panic(err.Error())
// 	}
// 	fmt.Println("db connected: ", &db)
// 	return db
// }

func Sss() string{
	 return "hello001"
}

// func dbInit(){
// 	 db := gormConnect();

// 		defer db.Close()
// 		db.Set("gorm:table_options", "ENGINE=InnoDB")
// 		db.AutoMigrate(&JsonRequest{})
// }

// func dbInsert(title string,disc string){
// 		db := gormConnect();

// 		defer db.Close()
// 		db.Create(&Card{Title:title, Description: disc})
// }

// func dbGetAll() []Card {
// 	db := gormConnect()

// 	defer db.Close()
// 	var cards []Card
// 	db.Order("created_at desc").Find(&cards)
// 	return cards
// }

// func IndexInit(c *gin.Context){
// 	dbInit()
// 	c.JSON(200, gin.H{"msg": "Init Db"})
// }

// func IndexGet(c *gin.Context){
// 	 cards := dbGetAll()
// 		c.JSON(200, gin.H{
// 			"cards": cards,
// 	 })
// }

// func IndexPost(c *gin.Context){
// 	 var json JsonRequest
// 		c.BindJSON(&json)

// 		// if err := c.ShouldBindJSON(&json); err != nil {
// 		// 	 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		// 		return
// 		// }
// 	 Title := c.PostForm("title")
// 	 Disc := c.PostForm("description")
// 		fmt.Println("@@@@@@@@@")
// 		log.Println(json.Description)
// 		log.Println(json.Title)
// 		dbInsert(Title,Disc)

// 		c.JSON(200, gin.H{
// 			"msg": "Hello world!!!!",
// 			"Title": Title,
// 			"log": "Hello world!!!!" + Disc,
// 	 })
// }
