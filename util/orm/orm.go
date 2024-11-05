package orm

import (
	str "github.com/Agelessbaby/BloomBlog/util/string"
	"reflect"
	"strings"
)

func GetGormFields(stc any) []string {
	typ := reflect.TypeOf(stc)
	if typ.Kind() == reflect.Ptr { //如果传的是指针类型，先解析指针
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Struct {
		columns := make([]string, 0, typ.NumField())
		for i := 0; i < typ.NumField(); i++ {
			fieldType := typ.Field(i)
			if fieldType.IsExported() { //只关注可导出成员
				if fieldType.Tag.Get("gorm") == "-" { //不做ORM映射的字段跳过
					continue
				}
				name := str.Camel2Snake(fieldType.Name) //如果没有gorm Tag，则把驼峰转为蛇形
				if len(fieldType.Tag.Get("gorm")) > 0 {
					content := fieldType.Tag.Get("gorm")
					if strings.HasPrefix(content, "column:") {
						content = content[7:]
						pos := strings.Index(content, ";")
						if pos > 0 {
							name = content[0:pos]
						} else if pos < 0 {
							name = content
						}
					}
				}
				columns = append(columns, name)
			}
		}
		return columns
	} else { //如果stc不是结构体则返回空切片
		return nil
	}
}
