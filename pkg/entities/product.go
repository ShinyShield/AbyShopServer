package entities

import "time"

type Product struct {
	ID            int32     `gorm:"primary_key"` // 產品的貨號
	Name          string    `gorm:"not null;"`   // 產品的名稱
	Price         int       `gorm:"not null;"`   // 產品的價錢
	ProductInfo   string    // 產品的資訊
	ImageURL      string    // 產品的圖片
	Owner         string    `gorm:"not null;"` // 產品的持有者
	MaxBuy        int32     `gorm:"not null;"` // 一次最多可以買的購買量
	ItemNumber    int32     `gorm:"not null;"` // 目前產品的剩餘數量
	AvailableTime time.Time // 什麼時候開放購買
}
