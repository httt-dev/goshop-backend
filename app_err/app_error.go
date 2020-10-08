package app_err

import "errors"

var (
	Http_AccessNotAllow = "Không cho phép truy cập"
	DataNotFound = errors.New("Không tồn tại dữ liệu")
	DataConflict = errors.New("Dữ liệu đã tồn tại")
)