package search

import (
	"errors"

	"github.com/ShinyShield/AbyShopServ/pkg/entities"
	"github.com/ShinyShield/AbyShopServ/pkg/store"
)

type Input struct {
	Query string
}

type Output struct {
	Products []*simpleProduct
	Count    int
}

type simpleProduct struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	ImageURL string `json:"image_url"`
}

func Exec(input Input) (Output, error) {

	var output Output

	query := store.DB.Model(entities.Product{}).Where("name LIKE ?", "%"+input.Query+"%")

	if err := query.Count(&output.Count).Error; err != nil {
		return output, errors.New("count error")
	}

	output.Products = make([]*simpleProduct, output.Count)

	if err := query.Scan(&output.Products).Error; err != nil {
		return output, errors.New("search error")
	}

	return output, nil
}
