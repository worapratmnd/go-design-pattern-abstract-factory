package models

type Nike struct {
}

type NikeShirt struct {
	Shirt
}

type NikeShoe struct {
	Shoe
}

func (n *Nike) MakeShirt() IShirt {
	shirt := Shirt{}
	shirt.SetLogo("Nike")
	shirt.SetSize(16)

	return &NikeShirt{Shirt: shirt}
}

func (n *Nike) MakeShoe() IShoe {
	shoe := Shoe{}
	shoe.SetLogo("Nike")
	shoe.SetSize(45)
	return &NikeShoe{Shoe: shoe}
}
