package dto

type SvcResponse struct {
	Error   bool        `json:"error"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
