package ginrestaurantliketransport

import (
	"be-food-delivery/common"
	"be-food-delivery/component"
	restaurantlikebiz "be-food-delivery/module/restaurantlike/biz"
	restaurantlikemodel "be-food-delivery/module/restaurantlike/model"
	restaurantlikestore "be-food-delivery/module/restaurantlike/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLikeRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestore.NewSQLStore(appCtx.GetMainDBConnection())
		//incStore := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, appCtx.GetPubSub())

		err = biz.UserLikeRestaurant(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
