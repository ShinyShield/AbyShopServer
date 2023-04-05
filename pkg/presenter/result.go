package presenter

// API 統一回傳結構
type Result struct {
	Data  interface{} `json:"data"`
	Code  StatusCode  `json:"code"` // 回應狀態碼
	Count int         `json:"count"`
}

type StatusCode uint16

// 回應狀態碼
const (
	StatusSuccess = 2001 // 成功

	StatusParamError   = 4001 // 參數有誤
	StatusTokenError   = 4011 // Token 有誤
	StatusAccountError = 4041 // 帳密認證有誤

	StatusServerError = 5001 // 伺服器錯誤
	StatusSQLError    = 5002 // 資料庫錯誤
)
