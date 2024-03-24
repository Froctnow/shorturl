package models

type CreateUrlRequest struct {
	URL string `json:"url" binding:"required"`
}

type CreateUrlResponse struct {
	Result string `json:"result"`
}
