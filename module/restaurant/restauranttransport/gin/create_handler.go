package ginrestaurant

import (
	"be-food-delivery/common"
	"be-food-delivery/component"
	"be-food-delivery/module/restaurant/restaurantbiz"
	"be-food-delivery/module/restaurant/restaurantmodel"
	"be-food-delivery/module/restaurant/restaurantstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.GenUID(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
