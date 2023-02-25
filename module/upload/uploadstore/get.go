package uploadstore

import (
	"be-food-delivery/common"
	"context"
)

func (store *sqlStore) GetImages(
	ctx context.Context,
	id int,
	moreKeys ...string,
) (*common.Image, error) {
	db := store.db
	var result common.Image

	db = db.Table(common.Image{}.TableName())

	if err := db.Where("id in (?)", id).First(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
