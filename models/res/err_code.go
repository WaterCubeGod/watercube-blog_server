package res

type CodeType int

const (
	SUCCESS       CodeType = 0
	Error         CodeType = 1
	SettingsError CodeType = 1001
)

var (
	ErrorMap = map[CodeType]string{
		SettingsError: "系统错误",
	}
)
