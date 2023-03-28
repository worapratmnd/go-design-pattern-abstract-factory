package models

type Adidas struct {
}

type AdidasShirt struct {
	Shirt
}

type AdidasShoe struct {
	Shoe
}

func (a *Adidas) MakeShirt() IShirt {
	shirt := Shirt{}
	shirt.SetLogo("Adidas")
	shirt.SetSize(14)

	return &AdidasShirt{Shirt: shirt}
}

func (a *Adidas) MakeShoe() IShoe {
	shoe := Shoe{}
	shoe.SetLogo("Adidas")
	shoe.SetSize(42)
	return &AdidasShoe{Shoe: shoe}
}
