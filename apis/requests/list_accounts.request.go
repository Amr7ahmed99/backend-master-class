package request_params

type ListAccountRequest struct {
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
}
