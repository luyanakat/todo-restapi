package ginhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/luyanakat/todo-restapi/common"
	"github.com/luyanakat/todo-restapi/modules/business"
	"github.com/luyanakat/todo-restapi/modules/item/model"
	"github.com/luyanakat/todo-restapi/modules/item/storage"
	"gorm.io/gorm"
	"net/http"
)

func ListItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var page common.Paging

		if err := c.ShouldBind(&page); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		page.Process()

		var filter model.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		store := storage.NewSQLStorage(db)
		bus := business.NewListItemBusiness(store)

		todoList, err := bus.ListItem(c.Request.Context(), &filter, &page)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(todoList, page, filter))
	}
}
