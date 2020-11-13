package utils

type APIError struct {
	Message    string `json:"err_message"`
	StatusCode int    `json:"status_code"`
}
