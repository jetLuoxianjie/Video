package defs


type Err struct {

	Error string `json:"error"`
	ErrorCode string `json:"error_code"`

}
type ErrResponse struct {
	HttpSC int
	Error Err
}

var(
	//通用错误 无法解析
	ErrorRequestBodyParseFailed=ErrResponse{400,Err{"Request body is not correct","001"}}
	//验证不通过没有这个用户
	ErrorNotAuthUser = ErrResponse{HttpSC: 401, Error: Err{Error: "User anthentication failed.", ErrorCode: "002"}}

)
