package ginrestaurant

import (
	"be-food-delivery/common"
	"be-food-delivery/component"
	"be-food-delivery/module/restaurant/restaurantbiz"
	"be-food-delivery/module/restaurant/restaurantstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteRestaurantHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		//id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
