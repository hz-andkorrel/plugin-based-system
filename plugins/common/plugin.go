package common

type Plugin interface {
	RegisterRoutes(router Router)
}
