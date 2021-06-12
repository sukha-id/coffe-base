package coffeDomain

type ListCoffe struct {
	Name string `json:"name"`
	Type string `json:"jenis"`
}

type DetailCoffe struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Stok        string `json:"stok"`
}

type InputCoffe struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
}
