package controller

type Result struct {
	Data   interface{} `json:"data"`
	Code   int         `json:"code"`
	ErrMsg string      `json:"err_msg"`
}

func NewResult(data interface{}, err error) Result {
	r := Result{}
	if err != nil {
		r.Code = 1
		r.ErrMsg = err.Error()
		return r
	}
	r.Data = data
	return r
}

func NewErrorResult(code int) Result {
	r := Result{}
	r.Code = code
	r.ErrMsg = errMsg[code]
	return r
}
func NewErrorDataResult(data interface{}, code int) Result {
	r := Result{}
	r.Code = code
	r.ErrMsg = errMsg[code]
	r.Data = data
	return r
}
