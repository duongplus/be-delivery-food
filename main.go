package main

import (
	"be-food-delivery/component"
	"be-food-delivery/component/uploadprovider"
	"be-food-delivery/middleware"
	"be-food-delivery/module/upload/uploadtransport/uploadgin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln(err)
	}

	appCtx := component.NewAppContext(db, s3Provider)

	if err = runService(appCtx); err != nil {
		log.Fatalln(err)
	}

}

func runService(appCtx component.AppContext) error {

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD

	r.POST("/upload", uploadgin.UploadHandler(appCtx))
	r.GET("upload/:id", uploadgin.GetHandler(appCtx))

	//restaurants := r.Group("/restaurants")
	//{
	//	restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
	//	restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
	//	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
	//	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
	//	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	//}

	return r.Run()
}
