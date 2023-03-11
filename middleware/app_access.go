package middleware

import (
	"be-food-delivery/common"
	"be-food-delivery/component"
	"github.com/gin-gonic/gin"
)

func AppAccess(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("api_key")
		if apiKey != common.ApiKey {
			panic(ErrWrongAuthHeader(nil))
		}

		c.Next()
	}
}
