package storage

import (
	"context"
	"github.com/luyanakat/todo-restapi/modules/item/model"
)

func (s *sqlStorage) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(model.TodoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status": "Deleted",
	}).Error; err != nil {
		return err
	}
	return nil
}
