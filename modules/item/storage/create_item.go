package storage

import (
	"context"

	"github.com/luyanakat/todo-restapi/modules/item/model"
)

func (s *sqlStorage) CreateItem(ctx context.Context, data *model.TodoItemCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
