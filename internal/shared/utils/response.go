package utils

type HttpResponse struct {
	Status  string `valid:"Required" json:"status"`
	Message string `valid:"Required" json:"message"`
}

type DefaultHttpResponse struct {
	HttpResponse
	Data interface{} `json:"data"`
}

func CreateHttpResponse(status string, message string, data interface{}) (response DefaultHttpResponse) {
	response = DefaultHttpResponse{
		HttpResponse: HttpResponse{
			Status:  status,
			Message: message,
		},
		Data: data,
	}

	return
}
