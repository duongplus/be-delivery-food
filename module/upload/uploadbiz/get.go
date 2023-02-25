package uploadbiz

import (
	"be-food-delivery/common"
	"context"
)

type GetImageStore interface {
	GetImages(
		ctx context.Context,
		id int,
		moreKeys ...string,
	) (*common.Image, error)
}

type getImageStoreBiz struct {
	store GetImageStore
}

func NewGetImageStoreBiz(store GetImageStore) *getImageStoreBiz {
	return &getImageStoreBiz{store: store}
}

func (biz *getImageStoreBiz) GetImages(
	ctx context.Context,
	id int,
) (*common.Image, error) {
	var result *common.Image

	result, err := biz.store.GetImages(ctx, id)

	if err != nil {
		return nil, err
	}

	return result, nil
}
