package subscriber

import (
	"be-food-delivery/common"
	"be-food-delivery/component"
	"be-food-delivery/module/restaurant/restaurantstore"
	"be-food-delivery/pubsub"
	"context"
)

type HasRestaurantId interface {
	GetRestaurantId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)

	store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData, ok := msg.Data().(HasRestaurantId)
			if ok {
				_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
			}
		}
	}()
}

func RunDecreaseLikeCountAfterUserUnlikeRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user unlikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
