package uploadmodel

import (
	"be-food-delivery/common"
)

const EntityName = "Upload"

type Upload struct {
	common.SQLModel `json:",inline"`
	common.Image    `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

//
//func (u *Upload) Mask(isAdmin bool) {
//	u.GenUID(common.DBTypeUpload, 1)
//}

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"file is not image",
		"ErrFileIsNotImage",
	)
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot save uploaded file",
		"ErrCannotSaveFile",
	)
}
