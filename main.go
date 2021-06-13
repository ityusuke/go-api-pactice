package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Age      int       `json:age`
	Birthday time.Time `json:birthday`
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "mymac"
	PASS := "mymac"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "GO_API"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open(DBMS, CONNECT + "?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
			panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	return db
}

func main() {
	r := gin.Default()

 r.GET("/hello", func(c *gin.Context) {
  c.String(http.StatusOK, "Hello world")
})
db := gormConnect()
db.Set("gorm:table_options", "ENGINE=InnoDB")
db.AutoMigrate(&User{})
defer db.Close()
db.LogMode(true)
	r.Run(":8080")
}
