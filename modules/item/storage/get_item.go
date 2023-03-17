package storage

import (
	"context"
	"github.com/luyanakat/todo-restapi/modules/item/model"
)

func (s *sqlStorage) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	if err := s.db.Table(model.TodoItem{}.TableName()).Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
