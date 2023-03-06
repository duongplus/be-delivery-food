package restaurantlikestore

import (
	"be-food-delivery/common"
	restaurantlikemodel "be-food-delivery/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) Find(ctx context.Context, userId, restaurantId int) (*restaurantlikemodel.Like, error) {
	var like restaurantlikemodel.Like
	db := s.db

	err := db.Where("restaurant_id = ? and user_id = ?", restaurantId, userId).
		First(&like).Error

	if err != nil {
		return nil, common.ErrDB(err)
	}

	return &like, nil
}
