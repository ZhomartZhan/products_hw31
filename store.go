package products_hw31

type ProductStore interface {
	Create(product *Product) (*Product, error)
	GetById(id string) (*Product, error)
	List() ([]Product, error)
	Delete(id string) error
}
