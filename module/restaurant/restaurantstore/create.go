package restaurantstore

import (
	"be-food-delivery/common"
	"be-food-delivery/module/restaurant/restaurantmodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
