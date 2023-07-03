package r

var (
	OK           = response(0, "成功")
	FAIL         = response(40001, "失败")
	PARAMS_ERROR = response(40002, "参数错误")
	SYSTEM_ERROR = response(40003, "系统错误")
	NO_AUTH      = response(40005, "无权限")
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func response(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code: OK.Code,
		Msg:  OK.Msg,
		Data: data,
	}
}

func (res *Response) WithMsg(msg string) *Response {
	return &Response{
		Code: res.Code,
		Msg:  msg,
		Data: res.Data,
	}
}
