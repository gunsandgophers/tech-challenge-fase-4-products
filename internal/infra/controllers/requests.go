package controllers

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (p *ProductRequest) ValidateProduct() error {
	if p.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if p.Category == "" {
		return errParamIsRequired("category", "string")
	}
	if p.Price == 0 {
		return errParamIsRequired("price", "float64")
	}
	return nil
}
