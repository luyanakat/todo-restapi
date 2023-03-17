package business

import (
	"context"
	"github.com/luyanakat/todo-restapi/modules/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

type getItemBusiness struct {
	store GetItemStorage
}

func NetGetItemBusiness(store GetItemStorage) *getItemBusiness {
	return &getItemBusiness{store: store}
}

func (b *getItemBusiness) GetItemByID(ctx context.Context, id int) (*model.TodoItem, error) {
	data, err := b.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
