package storage

import (
	"context"
	"github.com/luyanakat/todo-restapi/common"
	"github.com/luyanakat/todo-restapi/modules/item/model"
)

func (s *sqlStorage) ListItem(
	tx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TodoItem, error) {

	var todoList []model.TodoItem

	db := s.db.Where("status <> ?", "Deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id asc").Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&todoList).Error; err != nil {
		return nil, err
	}
	return todoList, nil
}
