package ocurrency

type IOccurrency interface {
	Store(*Ocurrency) error
	FindByID(id uint) (*Ocurrency, error)
	FindAll() ([]*Ocurrency, error)
}
