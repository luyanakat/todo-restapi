package ginhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/luyanakat/todo-restapi/common"
	"github.com/luyanakat/todo-restapi/modules/business"
	"github.com/luyanakat/todo-restapi/modules/item/model"
	"github.com/luyanakat/todo-restapi/modules/item/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		}

		store := storage.NewSQLStorage(db)
		bus := business.NewUpdateItemBusiness(store)

		err = bus.UpdateItemByID(c.Request.Context(), id, &data)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
