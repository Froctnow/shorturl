package models

type CreateURLRequest struct {
	URL string `json:"url" binding:"required"`
}

type CreateURLResponse struct {
	Result string `json:"result"`
}

type CreateBatchURLRequest struct {
	CorrelationID string `json:"correlation_id" binding:"required"`
	OriginalURL   string `json:"original_url" binding:"required"`
}

type CreateBatchURLResponse struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}
