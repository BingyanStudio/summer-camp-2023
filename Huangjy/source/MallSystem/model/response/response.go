package response

/*
 *	response的结构体
 */
type response struct {
	Succeed bool   `json:"succeed"`
	Error   string `json:"error"`
	Data    string `json:"data"`
}

/*
 *	生成一个response
 */
func MakeResponse(b bool, e string, d string) response {
	return response{
		Succeed: b,
		Error:   e,
		Data:    d,
	}
}

/*
 *	一个成功的response
 */
func MakeSucceedResponse(d string) response {
	return MakeResponse(true, "", d)
}

/*
 *	一个失败的response
 */
func MakeFailedResponse(e string) response {
	return MakeResponse(false, e, "")
}
