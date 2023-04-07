package buy_product

import (
	"errors"
	"time"

	"github.com/ShinyShield/AbyShopServ/pkg/entities"
	"github.com/ShinyShield/AbyShopServ/pkg/store"
)

type Input struct {
	ProductID  int32
	CustomerID int32
	Number     int32
}

func Exec(input Input) error {

	tx := store.DB.Begin()

	var product entities.Product

	query := tx.Model(entities.Product{}).
		Where("id = ?", input.ProductID).
		Scan(&product)
	if err := query.Error; err != nil {
		return err
	}

	currentTime := time.Now()
	if currentTime.Before(product.AvailableTime) {
		return errors.New("product time not available")
	}
	if input.Number > product.ItemNumber || input.Number > product.MaxBuy {
		return errors.New("product number not available")
	}

	update := tx.Model(entities.Product{}).
		Where("id = ?", input.ProductID).
		Update("item_number", (product.ItemNumber - input.Number))
	if err := update.Error; err != nil {
		return err
	}

	record := new(entities.Record)
	record.Number = input.Number
	record.CustomerID = input.CustomerID
	record.ProductID = input.ProductID
	record.Date = time.Now()

	add := tx.Model(entities.Record{}).Create(&record)
	if err := add.Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
