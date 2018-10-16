package e

var MsgFlags = map[int]string {
	SUCCESS: "ok",
	ERROR: "fail",
	INVAILD_PARAMS: "请求参数错误",
	ERROR_EXIST_TAG: "已存在该标签",
	ERROR_NOT_EXIST_TAG: "不存在的标签",
	ERROR_NOT_EXIST_ARTICLE: "不存在的文章",
	ERROR_AUTH_CHECK_TOKEN_FAIL: "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token鉴权超时",
	ERROR_AUTH_TOKEN: "Token生成失败",
	ERROR_AUTH: "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
