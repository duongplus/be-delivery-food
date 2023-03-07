package ginrestaurantliketransport

import (
	"be-food-delivery/common"
	"be-food-delivery/component"
	restaurantlikebiz "be-food-delivery/module/restaurantlike/biz"
	restaurantlikestore "be-food-delivery/module/restaurantlike/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserUnlikeRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestore.NewSQLStore(appCtx.GetMainDBConnection())
		//decStore := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.UserUnlikeRestaurantBiz(store, appCtx.GetPubSub())

		err = biz.UserLikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
