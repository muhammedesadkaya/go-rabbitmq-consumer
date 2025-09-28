package response

type JsonResponse struct {
	Code    bool        `json:"code"`
	Message string      `json:"message"`
	Count   int         `json:"count"`
	Data    interface{} `json:"data"`
}

func Response(code bool, message string, count int, data interface{}) *JsonResponse {
	return &JsonResponse{
		Code:    code,
		Message: message,
		Count:   count,
		Data:    data,
	}
}
