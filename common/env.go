package common

type EnvType int

const (
	Dev EnvType = iota
	Staging
	Production
)

func (EnvType) String(env EnvType) string {
	return []string{"Development", "Staging", "Production"}[env]
}

var CurrentEnv EnvType = Dev
