package common

type Plugin interface {
	Initialize(bus Bus)
}
