package restaurantlikebiz

import (
	"be-food-delivery/common"
	"be-food-delivery/component/asyncjob"
	model "be-food-delivery/module/restaurantlike/model"
	"context"
)

type (
	UserUnlikeRestaurantStore interface {
		Find(ctx context.Context, userId, restaurantId int) (*model.Like, error)
		Delete(ctx context.Context, userId, restaurantId int) error
	}
)

type DecreaseLikeRestaurantStore interface {
	DecreaseLikeCount(ctx context.Context, restaurantId int) error
}

type userUnlikeRestaurantBiz struct {
	store    UserUnlikeRestaurantStore
	decStore DecreaseLikeRestaurantStore
}

func UserUnlikeRestaurantBiz(store UserUnlikeRestaurantStore, decStore DecreaseLikeRestaurantStore) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{store: store, decStore: decStore}
}

func (biz *userUnlikeRestaurantBiz) UserLikeRestaurant(ctx context.Context, userId, restaurantId int) error {

	if _, err := biz.store.Find(ctx, userId, restaurantId); err != nil {
		return model.ErrCannotUnlikeRestaurant(err)
	}

	if err := biz.store.Delete(ctx, userId, restaurantId); err != nil {
		return model.ErrCannotUnlikeRestaurant(err)
	}

	go func() {
		common.AppRecover()

		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.decStore.DecreaseLikeCount(ctx, restaurantId)
		})

		group := asyncjob.NewGroup(true, job)

		_ = group.Run(ctx)
	}()

	return nil
}
