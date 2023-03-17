package business

import (
	"context"
	"errors"
	"strings"

	"github.com/luyanakat/todo-restapi/modules/item/model"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreate) error
}

type createItemBusiness struct {
	store CreateItemStorage
}

func NewCreateItemBusiness(store CreateItemStorage) *createItemBusiness {
	return &createItemBusiness{store: store}
}

func (b *createItemBusiness) CreateNewItem(ctx context.Context, data *model.TodoItemCreate) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return errors.New("title is blank")
	}
	if err := b.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
