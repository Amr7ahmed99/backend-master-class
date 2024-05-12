package requests

type GetUserRequest struct {
	Username string `json:"username" uri:"username" binding:"required".`
}
