package copy

import (
	"reflect"
)

// StructCopy copy field value from src to dest , same name ,same type value will copy
// 不需要返回值， 因为返回值是interface{}类型，调用端还需要做类型转换，更麻烦
func StructCopy(src interface{}, dst interface{}) {

	var vsrc reflect.Value
	var vdst reflect.Value
	if reflect.TypeOf(src).Kind() == reflect.Ptr {
		vsrc = reflect.ValueOf(src).Elem()
	} else {
		vsrc = reflect.ValueOf(src)
	}
	if reflect.TypeOf(dst).Kind() == reflect.Ptr {
		vdst = reflect.ValueOf(dst).Elem()
	} else {
		vdst = reflect.ValueOf(dst)
	}

	for i := 0; i < vsrc.NumField(); i++ {
		//非指针field call isnil会panic，所以更推荐用下面注释掉的check是否为zero，另外reflect还有个deepequal。
		dv := vdst.FieldByName(vsrc.Type().Field(i).Name)
		if !vsrc.Field(i).IsZero() && dv.IsValid() && dv.CanSet() {
			dv.Set(vsrc.Field(i))
		}

		//if src.Field(i).Interface() == reflect.Zero(src.Field(i).Type()).Interface() {
		//  fmt.Println("hehe", src.Type().Field(i).Type, src.Type().Field(i).Name)
		//} else {
		//  if dv := dst.Elem().FieldByName(src.Type().Field(i).Name); dv.IsValid() && dv.CanSet() {
		//      dv.Set(src.Field(i))
		//  }
		//
		//}
	}
}
