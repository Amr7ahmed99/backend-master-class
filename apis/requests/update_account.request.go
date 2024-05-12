package requests

type UpdateAccountRequest struct {
	Balance  int64 `json:"balance" binding:"required"`
	Currency int32 `json:"currency" binding:"required"`
}
