package reflectutil

import "reflect"

func IsPtr(obj interface{}) bool {
	return reflect.TypeOf(obj).Kind() == reflect.Ptr
}
