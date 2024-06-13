package inventory

type Receipt struct {
	Items []Product `json:"items"`
}

func (r *Receipt) AddProduct(p Product) {
	r.Items = append(r.Items, p)
}