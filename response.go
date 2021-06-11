package rin

type JSON map[string]interface{}

type IResponse interface {
	GetResponse() map[string]interface{}
	GetStatusCode() int
}

type BaseResponse struct {
	response JSON
	code int
}

func (r *BaseResponse) GetResponse() map[string]interface{} {
	return r.response
}

func (r *BaseResponse) GetStatusCode() int {
	return r.code
}

func Success(code int, response interface{}) IResponse {
	return &BaseResponse{code: code, response: JSON{
		"status":   "success",
		"response": response,
	}}
}

func Failed(code int, response interface{}) IResponse {
	return &BaseResponse{code: code, response: JSON{
		"status": "failed",
		"error":  response,
	}}
}