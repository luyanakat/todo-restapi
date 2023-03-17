package business

import (
	"context"
	"errors"
	"github.com/luyanakat/todo-restapi/modules/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
}

type updateItemBusiness struct {
	store UpdateItemStorage
}

func NewUpdateItemBusiness(store UpdateItemStorage) *updateItemBusiness {
	return &updateItemBusiness{store: store}
}

func (b *updateItemBusiness) UpdateItemByID(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {
	data, err := b.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if data.Status == "Deleted" {
		return errors.New("item is deleted")
	}

	if err := b.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}
	return nil
}
