package ginrestaurant

import (
	"be-food-delivery/common"
	"be-food-delivery/component"
	"be-food-delivery/module/restaurant/restaurantbiz"
	"be-food-delivery/module/restaurant/restaurantmodel"
	"be-food-delivery/module/restaurant/restaurantstore"
	restaurantlikestore "be-food-delivery/module/restaurantlike/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
		likeStore := restaurantlikestore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store, likeStore)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
