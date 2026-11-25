package types

type Category struct {
	ID   int64
	Name string
}

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       int64
	ImageURL    *string
	CategoryID  int64
	IsAvailable bool
}

type ProductInput struct {
	Name        string
	Description string
	Price       int64
	ImageURL    *string
	CategoryID  int64
	IsAvailable bool
}

// MenuSection is one category of a customer-facing menu with the items on sale in it.
type MenuSection struct {
	ID    int64
	Name  string
	Items []Product
}
