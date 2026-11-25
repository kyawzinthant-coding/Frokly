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
