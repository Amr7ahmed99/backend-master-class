package requests

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency int32  `json:"currency" binding:"required"`
}
