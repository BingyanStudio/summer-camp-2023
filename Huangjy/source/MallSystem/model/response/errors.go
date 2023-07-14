package response

/*
 *	定义若干常用的response，包括：
 *	未经许可的请求
 *	非法消息
 *	请求超时
 */
var (
	UnautherizedError = MakeFailedResponse("请先登录")
	InvalidInfoError  = MakeFailedResponse("不合法的请求")
	TimeoutError      = MakeFailedResponse("服务超时")
)
