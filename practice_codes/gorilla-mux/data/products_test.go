package data

import "testing"

func TestValidateFunc(t *testing.T) {
	p := &Product{
		ID:    123,
		Price: 1,
		Name:  "Rajesh",
		SKU:   "abs-abs-abs",
	}

	err := p.Validator()
	if err != nil {
		t.Errorf("should have failed, err=%s", err)
	}
}
