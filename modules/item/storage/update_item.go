package storage

import (
	"context"
	"github.com/luyanakat/todo-restapi/modules/item/model"
)

func (s *sqlStorage) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := s.db.Table(model.TodoItem{}.TableName()).Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
