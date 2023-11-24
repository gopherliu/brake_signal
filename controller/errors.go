package controller

const (
	INVALIDBLANCE = 2 + iota
)

var errMsg = map[int]string{
	INVALIDBLANCE: "无可用余额",
}
