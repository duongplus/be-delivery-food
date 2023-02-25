package uploadbiz

import (
	"be-food-delivery/common"
	"be-food-delivery/component/uploadprovider"
	"be-food-delivery/module/upload/uploadmodel"
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStore CreateImageStorage) *uploadBiz {
	return &uploadBiz{provider: provider, imgStore: imgStore}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                // "img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg

	dst := fmt.Sprintf("%s/%s", folder, fileName)

	img, err := biz.provider.SaveFileUploaded(ctx, data, dst)

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Key = dst
	img.Width = w
	img.Height = h
	//img.CloudName = "s3" // should be set in provider
	img.Extension = fileExt

	if err := biz.imgStore.CreateImage(ctx, img); err != nil {
		_ = biz.provider.DeleteFileUploaded(ctx, dst)
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
