package uploadgin

import (
	"be-food-delivery/common"
	"be-food-delivery/component"
	"be-food-delivery/module/upload/uploadbiz"
	"be-food-delivery/module/upload/uploadstore"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := uploadstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := uploadbiz.GetImageStore(store)

		data, err := biz.GetImages(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
