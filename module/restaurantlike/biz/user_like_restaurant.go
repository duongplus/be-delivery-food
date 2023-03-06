package restaurantlikebiz

import (
	"be-food-delivery/common"
	"be-food-delivery/component/asyncjob"
	restaurantlikemodel "be-food-delivery/module/restaurantlike/model"
	"context"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncreaseLikeRestaurantStore interface {
	IncreaseLikeCount(ctx context.Context, restaurantId int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncreaseLikeRestaurantStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, incStore IncreaseLikeRestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, incStore: incStore}
}

func (biz *userLikeRestaurantBiz) UserLikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {

	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	go func() {
		common.AppRecover()

		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
		})

		group := asyncjob.NewGroup(true, job)

		_ = group.Run(ctx)
	}()

	return nil
}
