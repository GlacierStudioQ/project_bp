package util

import (
	"strings"
	"reflect"
	"fmt"
)

//获取某结构体中某字段的某Tag的值，如果没有则返回空字符串
func GetStructTagName(structName interface{}, fieldName string, tagKey string) (tagVal string)  {  
    t := reflect.TypeOf(structName)  
    if t.Kind() == reflect.Ptr {  
        t = t.Elem()  
    }  
    if t.Kind() != reflect.Struct {  
        return ""  
	}
	
	// 遍历字段
    fieldNum := t.NumField()  
    for i := 0; i < fieldNum; i++ {
        fmt.Println("field name = " + t.Field(i).Name)
		if(0 == strings.Compare(fieldName, t.Field(i).Name)){
            tagVal = t.Field(i).Tag.Get(tagKey)
            return tagVal
		}
	}  
    return ""
}  