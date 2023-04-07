package entities

import "time"

type Product struct {
	ID            int32     `gorm:"primary_key"` //
	Name          string    `gorm:"not null;"`   //
	Price         int       `gorm:"not null;"`   //
	ProductInfo   string    //
	ImageURL      string    //
	Owner         string    `gorm:"not null;"` //
	MaxBuy        int32     `gorm:"not null;"` //
	ItemNumber    int32     `gorm:"not null;"` //
	AvailableTime time.Time //
}
