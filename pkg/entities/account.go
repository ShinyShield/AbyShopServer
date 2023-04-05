package entities

type Account struct {
	ID       int32  `gorm:"primary_key"` // 帳戶 ID
	Email    string `gorm:"not null"`    // 帳戶 email
	Name     string `gorm:"not null"`    // 帳戶號碼
	Password string `gorm:"not null;"`   // 帳戶密碼
	Token    string `gorm:"unique;"`     // 帳戶 Token
}
