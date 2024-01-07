package request_params

type ListUsersRequest struct {
	PageSize int32 `json:"page_size" binding:"required,min=5,max=10"`
	PageID   int32 `json:"page_id" binding:"required,min=1"`
}
