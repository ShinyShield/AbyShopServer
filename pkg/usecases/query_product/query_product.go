package query_product

import (
	"errors"
	"fmt"

	"github.com/ShinyShield/AbyShop-server/pkg/entities"
	"github.com/ShinyShield/AbyShop-server/pkg/store"
)

type Input struct {
	ProductID int32
}

type Output struct {
	Product SimpleProduct
}

type SimpleProduct struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	ProductInfo string `json:"product_info"`
	ImageURL    string `json:"image_url"`
	MaxBuy      int32  `json:"max_buy"`
	ItemNumber  int32  `json:"item_number"`
}

func Exec(input Input) (Output, error) {

	var s SimpleProduct
	query := store.DB.Model(entities.Product{}).Where("id = ?", input.ProductID).
		Scan(&s)
	if err := query.Error; err != nil {
		fmt.Println(err.Error())
		return Output{}, errors.New("product query error")
	}

	var output Output
	output.Product = s

	return output, nil
}
