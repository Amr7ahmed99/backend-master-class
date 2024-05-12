package request_params

type CreateTransferRequest struct {
	Amount      int64 `json:"amount" binding:"required,gt=0"`
	Currency    int32 `json:"currency" binding:"required,min=1"`
	FromAccount int64 `json:"from_account" binding:"required,min=1"`
	ToAccount   int64 `json:"to_amount" binding:"required,min=1"`
}
