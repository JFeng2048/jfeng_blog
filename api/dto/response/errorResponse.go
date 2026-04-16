package response

// ErrorResponse 错误响应体
type ErrorResponse struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误信息
	Data interface{} `json:"data"` // 错误数据
}
