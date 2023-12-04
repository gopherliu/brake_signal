package controller

const (
	INVALIDBLANCE = 2 + iota
	INVALIDVIN
	INVALIDTIME
)

var errMsg = map[int]string{
	INVALIDBLANCE: "无可用余额",
	INVALIDVIN:    "无效VIN",
	INVALIDTIME:   "无效时间",
}
