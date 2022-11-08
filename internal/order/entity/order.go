package entity

import "errors"

type Order struct {
	ID                     string
	Price, Tax, FinalPrice float64
}

func NewOrder(id string, price, tax float64) (*Order, error) {
	order := &Order{
		ID:         id,
		Price:      price,
		Tax:        tax,
		FinalPrice: price + tax,
	}
	if err := order.isValid(); err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) isValid() error {
	if o.ID == "" {
		return errors.New("invalid id")
	}
	if o.Price <= 0 {
		return errors.New("invalid price")
	}
	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func (o *Order) CalculatePrice() error {
	o.FinalPrice = o.Price + o.Tax
	if err := o.isValid(); err != nil {
		return err
	}
	return nil
}
