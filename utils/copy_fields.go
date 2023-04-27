package utils

import (
	"reflect"
)

// CopyFields 用b的所有字段覆盖a的
// CopyFields 如果fields不为空, 表示用b的特定字段覆盖a的
// CopyFields a,b应该为结构体指针
func CopyFields(a interface{}, b interface{}, fields ...string) {
	//at := reflect.TypeOf(a)
	av := reflect.ValueOf(a)
	bt := reflect.TypeOf(b)
	bv := reflect.ValueOf(b)

	// 要复制哪些字段
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < bv.Elem().NumField(); i++ {
			_fields = append(_fields, bt.Elem().Field(i).Name)
		}
	}

	// 复制
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		f := av.Elem().FieldByName(name)
		bValue := bv.Elem().FieldByName(name)

		// a中有同名的字段并且类型一致才复制
		if f.CanSet() && f.IsValid() && f.Kind() == bValue.Kind() {
			f.Set(bValue)
		}
	}
	return
}
