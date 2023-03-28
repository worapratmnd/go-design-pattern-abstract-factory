package models

import "errors"

type IFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func GetSportFactory(brand string) (IFactory, error) {
	var (
		result IFactory
	)
	if brand == "adidas" {
		result = &Adidas{}
		return result, nil
	} else if brand == "nike" {
		result = &Nike{}
		return result, nil
	}

	return nil, errors.New("WRONG BRAND TYPE PASSED")
}
