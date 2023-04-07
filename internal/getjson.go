package getjson

// TODO: output error
// TODO: get into awesome-go
// pkg.go.dev
// goreportcard.com
// coverage

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/limiu82214/getjson/pkg/gojpath"
)

var unmarshal func([]byte, interface{}) error

func Unmarshal(jsonString []byte, e interface{}) error {
	realData, err := getReadData(jsonString)
	if err != nil {
		return err
	}

	return mapIt(realData, e)
}

func mapIt(realData, e interface{}) error {
	rv := reflect.ValueOf(e)
	if !isPtrOfStruct(rv) {
		return fmt.Errorf("e must be a ptr of struct")
	}

	rt := reflect.TypeOf(e)
	eFieldCount := rt.Elem().NumField()

	for i := 0; i < eFieldCount; i++ {
		etField := rt.Elem().Field(i)
		jpath := etField.Tag.Get("getjson")
		field := rv.Elem().Field(i)

		switch {
		case jpath == "" && isStruct(field):
			err := mapIt(realData, field.Addr().Interface())
			if err != nil {
				return err
			}
		case jpath == "" && isPtrOfStruct(field):
			err := mapIt(realData, field.Interface())
			if err != nil {
				return err
			}
		default:
			v := gojpath.Get(realData, jpath)
			realDataRV := reflect.ValueOf(v)

			realDataRV, err := changeRealDataType(realDataRV, etField.Type)
			if err != nil {
				return err
			}

			setDataToField(field, realDataRV)
		}
	}

	return nil
}

func SetUnmarshalFunc(f func([]byte, interface{}) error) {
	unmarshal = f
}

func callUnmarshal(jsonString []byte, data interface{}) error {
	var err error
	if unmarshal != nil {
		err = unmarshal(jsonString, &data)
	} else {
		err = json.Unmarshal(jsonString, &data)
	}

	if err != nil {
		return fmt.Errorf("json.Unmarshal error: %v", err)
	}

	return nil
}

func getReadData(jsonString []byte) (interface{}, error) {
	var data interface{}

	err := callUnmarshal(jsonString, &data)
	if err != nil {
		return nil, fmt.Errorf("callUnmarshal error: %v", err)
	}

	return data, nil
}

func valid(rv reflect.Value) error {
	if isPtrOfStruct(rv) {
		return fmt.Errorf("e must be a ptr of struct")
	}

	return nil
}

func isStruct(rv reflect.Value) bool {
	return rv.Kind() == reflect.Struct
}

func isPtrOfStruct(rv reflect.Value) bool {
	if rv.Kind() == reflect.Ptr {
		if rv.Elem().Kind() == reflect.Struct {
			return true
		}
	}

	return false
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
