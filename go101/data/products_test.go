package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "Latte",
		Price: 1.2,
		SKU:   "anb-asd-adfdf",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
