package funcs

import (
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

func StructToMap(obj interface{}) map[string]interface{} {
	v := reflect.Indirect(reflect.ValueOf(obj))
	t := v.Type()
	if t.Kind() != reflect.Struct {
		err := errors.New("obj must be a struct or a pointer to a struct")
		panic(err)
	}
	resultMap := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue // Skip fields without json tag or with "-" tag
		}
		resultMap[tag] = fieldValue.Interface()
	}
	return resultMap
}

//Struct2MapString 结构体转json
//使用场景 resty curl 构造请求参数时常用
func Struct2MapString(i interface{}) (out map[string]string, err error) {

	var myMap map[string]interface{}
	data, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &myMap)
	if err != nil {
		return nil, err
	}
	out = make(map[string]string)
	for k, v := range myMap {
		out[k] = cast.ToString(v)
	}
	return out, nil
}
