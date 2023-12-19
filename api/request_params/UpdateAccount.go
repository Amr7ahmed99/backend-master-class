package request_params

type UpdateAccountRequest struct {
	Balance  int64  `json:"balance" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}
