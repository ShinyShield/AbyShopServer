package logout

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/ShinyShield/AbyShop-server/pkg/entities"
	"github.com/ShinyShield/AbyShop-server/pkg/store"
)

type Input struct {
	AccountID int32
}

func Exec(input Input) error {

	update := store.DB.Model(entities.Account{}).
		Where("id = ?", input.AccountID).
		Update("token", gorm.Expr("NULL"))
	if err := update.Error; err != nil {
		return errors.New("update error")
	}

	return nil
}
