package res

type CodeType int

const (
	SUCCESS       CodeType = 0
	Error         CodeType = 1
	SettingsError CodeType = 1001 // 系统错误
	ArgumentError CodeType = 1002 // 参数错误
	UploadError   CodeType = 1003 // 上传错误
)

var (
	ErrorMap = map[CodeType]string{
		SettingsError: "系统错误",
		ArgumentError: "参数错误",
		UploadError:   "上传错误",
	}
)
