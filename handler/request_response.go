package handler

type FormatRequest struct {
	Zip     bool   `json:"zip"`
	Sep     string `json:"sep"`
	Content string `json:"content"`
}

type FormatResponse struct {
	Content string `json:"content"`
}
