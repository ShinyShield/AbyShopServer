package entities

type Account struct {
	ID       int32  `gorm:"primary_key"` //  ID
	Email    string `gorm:"not null"`    //  email
	Name     string `gorm:"not null"`    //
	Password string `gorm:"not null;"`   //
	Token    string `gorm:"unique;"`     //  Token
}
