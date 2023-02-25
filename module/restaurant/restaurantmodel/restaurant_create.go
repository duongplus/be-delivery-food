package restaurantmodel

import (
	"be-food-delivery/common"
	"strings"
)

type RestaurantCreate struct {
	Id    int            `json:"id" gorm:"column:id;"`
	Name  string         `json:"name" gorm:"column:name;"`
	Addr  string         `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return common.ErrNameCannotBeEmpty
	}

	return nil
}
