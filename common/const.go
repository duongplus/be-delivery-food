package common

const (
	DbTypeRestaurant = 1
	DbTypeFood       = 2
	DbTypeCategory   = 3
	DBTypeUpload     = 4
	DbTypeUser       = 5
)

const CurrentUser = "user"
const ApiKey = "0x54519e9E7fB050b7a51E7fE38c1dc13222F9dE17"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
