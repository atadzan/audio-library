package v1

// error messages
const (
	errInvalidInputParamsMsg = "invalid input params"
	errInternalServerMsg     = "internal server error"
	errUserNotFoundMsg       = "user not found"
	errOperationFailedMsg    = "operation failed"
)

// success messages
const (
	successMsg = "success"
)

type Response struct {
	Message string `json:"message"`
}

type AuthResp struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func newMsgResponse(msg string) Response {
	return Response{Message: msg}
}

func newSuccessAuthResp(token string) AuthResp {
	return AuthResp{
		Message: successMsg,
		Token:   token,
	}
}
