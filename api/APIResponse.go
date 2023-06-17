package api

type APIResponse struct {
	ErrorCode    int    `json:"code"`
	ErrorMessage string `json:"message"`
	Result       string `json:"result"`
}
