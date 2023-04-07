package presenter

// API
type Result struct {
	Data  interface{} `json:"data"`
	Code  StatusCode  `json:"code"` //
	Count int         `json:"count"`
}

type StatusCode uint16

const (
	StatusSuccess = 2001 //

	StatusParamError   = 4001 //
	StatusTokenError   = 4011 // Token
	StatusAccountError = 4041 //

	StatusServerError = 5001 //
	StatusSQLError    = 5002 //
)
