package entities

import "time"

type Record struct {
	ID         int32     `gorm:"primary_key"` //
	ProductID  int32     //
	CustomerID int32     //
	Number     int32     //
	Date       time.Time //
}
