package ctype

import "encoding/json"

type ImageType int

const (
	Local     ImageType = 1 // 本地
	Telegraph ImageType = 2 // telegraph
)

func (s ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() string {
	switch s {
	case Local:
		return "本地"
	case Telegraph:
		return "telegraph"
	default:
		return "其他"
	}
}
