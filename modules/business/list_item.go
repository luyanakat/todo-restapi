package business

import (
	"context"
	"github.com/luyanakat/todo-restapi/common"
	"github.com/luyanakat/todo-restapi/modules/item/model"
)

type ListItemStorage interface {
	ListItem(
		tx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type listItemBusiness struct {
	store ListItemStorage
}

func NewListItemBusiness(store ListItemStorage) *listItemBusiness {
	return &listItemBusiness{store: store}
}

func (b *listItemBusiness) ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.TodoItem, error) {
	data, err := b.store.ListItem(ctx, filter, paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}
