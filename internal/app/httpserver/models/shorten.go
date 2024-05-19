package models

type CreateURLRequest struct {
	URL string `json:"url" binding:"required"`
}

type CreateURLResponse struct {
	Result string `json:"result"`
}
