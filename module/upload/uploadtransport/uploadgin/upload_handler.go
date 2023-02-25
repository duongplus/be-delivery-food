package uploadgin

import (
	"be-food-delivery/common"
	"be-food-delivery/component"
	"be-food-delivery/module/upload/uploadbiz"
	"be-food-delivery/module/upload/uploadstore"
	"github.com/gin-gonic/gin"
	_ "image/jpeg"
	_ "image/png"
)

func UploadHandler(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		//db := appCtx.GetMainDBConnection()

		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // we can close here

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		imgStore := uploadstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider(), imgStore)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
