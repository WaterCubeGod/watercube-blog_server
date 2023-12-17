package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// GetValidMsg 返回结构体中的msg参数
func GetValidMsg(err error, obj any, field string) string {
	getObj := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误消息
			// 根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get(field)
				return msg
			}
		}
	}
	return err.Error()
}
