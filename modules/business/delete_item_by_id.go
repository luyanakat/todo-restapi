package business

import (
	"context"
	"errors"
	"github.com/luyanakat/todo-restapi/modules/item/model"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type deleteItemBusiness struct {
	store DeleteItemStorage
}

func NewDeleteItemBusiness(store DeleteItemStorage) *deleteItemBusiness {
	return &deleteItemBusiness{store: store}
}

func (b *deleteItemBusiness) DeleteItemByID(ctx context.Context, id int) error {
	data, err := b.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if data.Status == "Deleted" {
		return errors.New("item was deleted")
	}

	if err := b.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
