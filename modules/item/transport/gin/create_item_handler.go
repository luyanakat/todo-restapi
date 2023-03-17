package ginhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/luyanakat/todo-restapi/modules/business"
	"github.com/luyanakat/todo-restapi/modules/item/model"
	"github.com/luyanakat/todo-restapi/modules/item/storage"
	"gorm.io/gorm"
	"net/http"
)

func CreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.TodoItemCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStorage(db)
		bus := business.NewCreateItemBusiness(store)

		err := bus.CreateNewItem(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data.ID,
		})
	}
}
