package response

// 返回结果结构
type DefaultHandlerResponse struct {
	Code    int         `json:"code" desc:"返回码"`
	Data    interface{} `json:"data" desc:"返回数据"`
	Message string      `json:"message" desc:"返回消息（比如Error信息）"`
}

func NewDefaultHandlerResponse(resp interface{}, err error) *DefaultHandlerResponse {
	if err != nil {
		return &DefaultHandlerResponse{
			Code:    CODE_FAIL,
			Data:    resp,
			Message: err.Error(),
		}
	}
	return &DefaultHandlerResponse{
		Code:    CODE_SUCCESS,
		Data:    resp,
		Message: MESSAGE_SUCCESS,
	}
}
