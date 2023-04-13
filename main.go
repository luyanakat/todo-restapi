package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	ginhandler "github.com/luyanakat/todo-restapi/modules/item/transport/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(db)

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginhandler.CreateItem(db))
			items.GET("", ginhandler.ListItem(db))
			items.GET("/:id", ginhandler.GetItem(db))
			items.PATCH("/:id", ginhandler.UpdateItem(db))
			items.DELETE("/:id", ginhandler.DeleteItem(db))
		}
	}

	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
