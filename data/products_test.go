package products

import "testing"

func TestValidateProduct(t *testing.T) {
	p := &Product{
		Name: "Pranav",
		Price: 1.00,
		SKU: "abc-abc-abc",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}