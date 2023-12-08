package app

type APIPorts interface {
	GetInfo(id string) error
}
