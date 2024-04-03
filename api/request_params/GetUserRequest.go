package request_params

type GetUserRequest struct {
	Username string `json:"username" uri:"username" binding:"required".`
}
