package registration

import (
	"errors"
	"fmt"
	"github.com/ShinyShield/AbyShopServ/pkg/entities"
	"github.com/ShinyShield/AbyShopServ/pkg/store"
	"github.com/ShinyShield/AbyShopServ/pkg/util"
)

type Input struct {
	Email    string
	Password string
}

type OutPut struct {
	Token string `json:"token"`
	ID    int32  `json:"id"`
	Name  string `json:"name"`
}

func Register(input Input) (OutPut, error) {

	var output OutPut
	var account entities.Account
	var count int
	// check if email already exists in database
	check := store.DB.Model(entities.Account{}).
		Where("email = ?", input.Email).
		Count(&count)
	if err := check.Error; err != nil {
		fmt.Println(err.Error())
		return output, errors.New("database error")
	}
	if count > 0 {
		return output, errors.New("email already exists")
	}

	// create new account
	account = entities.Account{
		//Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := store.DB.Create(&account).Error; err != nil {
		fmt.Println(err.Error())
		return output, errors.New("database error")
	}

	// generate token
	token, err := util.CreateToken(input.Email, input.Password)
	if err != nil {
		fmt.Println(err.Error())
		return output, errors.New("create token error")
	}

	// update account with token
	update := store.DB.Model(entities.Account{}).
		Where("email = ?", input.Email).
		Update("token", token)
	if err := update.Error; err != nil {
		fmt.Println(err.Error())
		return output, errors.New("update token error")
	}

	output.Token = token
	output.ID = account.ID
	output.Name = account.Name

	return output, nil
}

/*func Exec(input Input) (OutPut, error) {

	var output OutPut
	var account entities.Account

	query := store.DB.Model(entities.Account{}).
		Where("email = ? AND password = ?", input.Email, input.Password).
		Limit(1).
		Scan(&account)
	if err := query.Error; err != nil {
		fmt.Println(err.Error())
		return output, errors.New("can't find account")
	}

	// token
	token, err := util.CreateToken(input.Email, input.Password)
	if err != nil {
		fmt.Println(err.Error())
		return output, errors.New("create token error")
	}

	update := store.DB.Model(entities.Account{}).
		Where("email = ? AND password = ?", input.Email, input.Password).
		Update("token", token)
	if err := update.Error; err != nil {
		fmt.Println(err.Error())
		return output, errors.New("update token error")
	}

	output.Token = token
	output.ID = account.ID
	output.Name = account.Name

	return output, nil
}*/
