package getjson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func Unmarshal(jsonString []byte, e interface{}) error {
	realData, err := getReadData(jsonString)
	if err != nil {
		return err
	}

	rv := reflect.ValueOf(e)
	if err2 := valid(rv); err2 != nil {
		return err2
	}

	rt := reflect.TypeOf(e)
	eFieldCount := rt.Elem().NumField()

	for i := 0; i < eFieldCount; i++ {
		etField := rt.Elem().Field(i)
		jpath := etField.Tag.Get("getjson")

		// FIXME: support json path
		// FIXME: support nested json path

		realDataRV := getDataByJPath(realData, jpath)

		realDataRV, err = changeRealDataType(realDataRV, etField.Type)
		if err != nil {
			return err
		}

		f := rv.Elem().Field(i)
		setDataToField(f, realDataRV)
	}

	return nil
}

func getReadData(jsonString []byte) (map[string]interface{}, error) {
	realData := make(map[string]interface{})
	err := json.Unmarshal(jsonString, &realData)

	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %v", err)
	}

	return realData, nil
}

func valid(rv reflect.Value) error {
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("e must be a ptr")
	}

	if rv.Elem().Kind() != reflect.Struct { // ?
		return fmt.Errorf("e must be a ptr of struct")
	}

	return nil
}

func getDataByJPath(realData map[string]interface{}, jpath string) reflect.Value {
	return reflect.ValueOf(realData[jpath])
}

func changeRealDataType(realDataRV reflect.Value, targetType reflect.Type) (reflect.Value, error) {
	if targetType.String() == "time.Time" {
		tmpTime, err := time.Parse(time.RFC3339, realDataRV.String())
		if err != nil {
			return reflect.Value{}, fmt.Errorf("time.Parse error: %v", err)
		}

		realDataRV = reflect.ValueOf(tmpTime)
	}

	return realDataRV, nil
}

func setDataToField(field, realDataRV reflect.Value) {
	if field.IsValid() {
		if field.CanSet() {
			if field.Kind() == realDataRV.Kind() {
				field.Set(realDataRV)
			}
		}
	}
}
