package response

/*
 *	定义若干常用的response，包括：
 *	未经许可的请求
 *	非法消息
 *	请求超时
 */
var (
	UnautherizedError = MakeFailedResponse("Please login first")
	InvalidInfoError  = MakeFailedResponse("Invalid Info")
	TimeoutError      = MakeFailedResponse("Timeout")
)
