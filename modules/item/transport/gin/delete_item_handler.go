package ginhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/luyanakat/todo-restapi/common"
	"github.com/luyanakat/todo-restapi/modules/business"
	"github.com/luyanakat/todo-restapi/modules/item/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStorage(db)
		bus := business.NewDeleteItemBusiness(store)

		if err := bus.DeleteItemByID(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
