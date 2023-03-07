package restaurantlikebiz

import (
	"be-food-delivery/common"
	model "be-food-delivery/module/restaurantlike/model"
	"be-food-delivery/pubsub"
	"context"
)

type (
	UserUnlikeRestaurantStore interface {
		Find(ctx context.Context, userId, restaurantId int) (*model.Like, error)
		Delete(ctx context.Context, userId, restaurantId int) error
	}
)

//type DecreaseLikeRestaurantStore interface {
//	DecreaseLikeCount(ctx context.Context, restaurantId int) error
//}

type userUnlikeRestaurantBiz struct {
	store UserUnlikeRestaurantStore
	//decStore DecreaseLikeRestaurantStore
	pubSub pubsub.Pubsub
}

func UserUnlikeRestaurantBiz(
	store UserUnlikeRestaurantStore,
	//decStore DecreaseLikeRestaurantStore,
	pubSub pubsub.Pubsub,
) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{
		store: store,
		//decStore: decStore,
		pubSub: pubSub,
	}
}

func (biz *userUnlikeRestaurantBiz) UserLikeRestaurant(ctx context.Context, userId, restaurantId int) error {

	if _, err := biz.store.Find(ctx, userId, restaurantId); err != nil {
		return model.ErrCannotUnlikeRestaurant(err)
	}

	if err := biz.store.Delete(ctx, userId, restaurantId); err != nil {
		return model.ErrCannotUnlikeRestaurant(err)
	}

	_ = biz.pubSub.Publish(
		ctx,
		common.TopicUserDislikeRestaurant,
		pubsub.NewMessage(&model.Like{
			RestaurantId: restaurantId,
			UserId:       userId,
		}),
	)

	//go func() {
	//	common.AppRecover()
	//
	//	job := asyncjob.NewJob(func(ctx context.Context) error {
	//		return biz.decStore.DecreaseLikeCount(ctx, restaurantId)
	//	})
	//
	//	group := asyncjob.NewGroup(true, job)
	//
	//	_ = group.Run(ctx)
	//}()

	return nil
}
