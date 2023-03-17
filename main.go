package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginhandler "github.com/luyanakat/todo-restapi/modules/item/transport/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:abc@123@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
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

	err = r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
