package component

import (
	"be-food-delivery/component/uploadprovider"
	"be-food-delivery/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.Pubsub
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secretKey  string
	pb         pubsub.Pubsub
}

func NewAppContext(
	db *gorm.DB,
	upProvider uploadprovider.UploadProvider,
	secretKey string,
	pb pubsub.Pubsub,
) *appCtx {
	return &appCtx{
		db:         db,
		upProvider: upProvider,
		secretKey:  secretKey,
		pb:         pb,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB { return ctx.db }

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider { return ctx.upProvider }

func (ctx *appCtx) SecretKey() string { return ctx.secretKey }

func (ctx *appCtx) GetPubSub() pubsub.Pubsub { return ctx.pb }
