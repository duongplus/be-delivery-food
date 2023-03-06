package restaurantlikestore

import (
	"be-food-delivery/common"
	restaurantlikemodel "be-food-delivery/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
