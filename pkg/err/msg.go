package err

var CodeMsgMap = map[int]string{
	SUCCESS: "ok",
	SYSTEM_ERROR: "system error",
	INVALID_PARAM: "param error",
	UNAUTHORIZED: "Unauthorized",
	NOTFOUND: "not found",

}

func GetCodeMsg(code int) string {
	if msg, ok := CodeMsgMap[code]; ok {
		return msg
	}

	return CodeMsgMap[SYSTEM_ERROR]
}