package entity

import (
	"errors"
	"time"

	"github.com/Alefeoliveira/goexpert/api/pkg/entity"
)

var (
	errIDIsRequired    = errors.New("id is required")
	errInvalidID       = errors.New("invalid id")
	errNameIsRequired  = errors.New("name is required")
	errPriceIsRequired = errors.New("price is required")
	errInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return errIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return errInvalidID
	}
	if p.Name == "" {
		return errNameIsRequired
	}
	if p.Price == 0 {
		return errPriceIsRequired
	}
	if p.Price <= 0 {
		return errInvalidPrice
	}
	return nil
}
