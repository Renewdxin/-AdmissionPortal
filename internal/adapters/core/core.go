package core

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (user *Adapter) Create(id string) error {
	return nil
}

func (user *Adapter) Find(id string) error {
	return nil
}

func (user *Adapter) Update(id string) error {
	return nil
}

func (user *Adapter) Delete(id string) error {
	return nil
}

func (user *Adapter) FindByEmail(id string) error {
	return nil
}
