package response

// SuccessResponse 成功响应体
type SuccessResponse struct {
	Code int         `json:"code"` // 成功码
	Msg  string      `json:"msg"`  // 成功信息
	Data interface{} `json:"data"` // 成功数据
}