package ports

type APIPorts interface {
	GetInfo(id string) error
}
