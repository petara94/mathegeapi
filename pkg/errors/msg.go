package errors

var MsgFlags = map[int]string{
	SUCCESS:              "ok",
	ERROR:                "fail",
	INVALID_PARAMS:       "bad parameters",
	ERROR_NOT_EXIST:      "not exists",
	ERROR_NOT_EXIST_TASK: "task not exists",
	ERROR_WRONG_JSON:     "bad json data",
	ERROR_ADD:            "can't add data",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
