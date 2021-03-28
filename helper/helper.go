package helper

type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func APIResponse(message string, status int, data interface{}) Response {

	jsonResponse := Response{
		Message: message,
		Status:  status,
		Data:    data,
	}

	return jsonResponse
}
