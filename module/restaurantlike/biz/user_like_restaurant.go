package restaurantlikebiz

import (
	"be-food-delivery/common"
	restaurantlikemodel "be-food-delivery/module/restaurantlike/model"
	"be-food-delivery/pubsub"
	"context"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

//type IncreaseLikeRestaurantStore interface {
//	IncreaseLikeCount(ctx context.Context, restaurantId int) error
//}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	//incStore IncreaseLikeRestaurantStore
	pubSub pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	//incStore IncreaseLikeRestaurantStore,
	pubSub pubsub.Pubsub,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
		//incStore: incStore,
		pubSub: pubSub,
	}
}

func (biz *userLikeRestaurantBiz) UserLikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {

	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	_ = biz.pubSub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))

	//go func() {
	//	common.AppRecover()
	//
	//	job := asyncjob.NewJob(func(ctx context.Context) error {
	//		return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//	})
	//
	//	group := asyncjob.NewGroup(true, job)
	//
	//	_ = group.Run(ctx)
	//}()

	return nil
}
