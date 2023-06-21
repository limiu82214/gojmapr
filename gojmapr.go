// Package gojmapr is a package for mapping json data to struct.
//
// Package gojmapr 是一個用來將 json 資料映射到 struct 的套件。
package gojmapr

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/limiu82214/gojpath"
)

var unmarshalFunc func([]byte, interface{}) error //nolint:gochecknoglobals // I want user can use other json plugin

// ErrInputNotPtrOfStruct is a error for input is not a ptr of struct.
//
// ErrInputNotPtrOfStruct 是一個錯誤，用來表示輸入的參數不是一個指向 struct 的指標。
var ErrInputNotPtrOfStruct = fmt.Errorf("input must be a ptr of struct")

// ErrMapDataToNilPtr is a error for mapping data to nil ptr.
//
// ErrMapDataToNilPtr 是一個錯誤，用來表示映射資料到 nil 指標。
var ErrMapDataToNilPtr = fmt.Errorf("can not map data to nil ptr")

// Unmarshal is a function for mapping json data to struct use like json.Unmarshal().
//
// Unmarshal 是一個用來將 json 資料映射到 struct 的函式，使用方式與 json.Unmarshal() 相同。
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
		return ErrInputNotPtrOfStruct
	}

	rt := reflect.TypeOf(e)
	eFieldCount := rt.Elem().NumField()

	// Loop all field of struct
	// If field is a struct, call mapIt() again.
	// Otherwise use gojpath to get data from json.
	for i := 0; i < eFieldCount; i++ {
		tField := rt.Elem().Field(i)
		jpath := tField.Tag.Get("gojmapr")
		vField := rv.Elem().Field(i)

		switch {
		case jpath == "" && isStruct(vField):
			err := mapIt(realData, vField.Addr().Interface())
			if err != nil {
				return err
			}
		case jpath == "" && !isPtrNil(vField) && isPtrOfStruct(vField):
			err := mapIt(realData, vField.Interface())
			if err != nil {
				return err
			}
		case jpath == "" && isPtrNil(vField):
			return ErrMapDataToNilPtr
		case jpath == "":
			continue
		default:
			v, err := gojpath.Get(realData, jpath)
			if err != nil {
				return fmt.Errorf("gojpath.Get error: %w", err)
			}

			realDataRV := reflect.ValueOf(v)

			realDataRV, err = changeRealDataType(realDataRV, tField.Type)
			if err != nil {
				return err
			}

			setDataToField(vField, realDataRV)
		}
	}

	return nil
}

// SetUnmarshalFunc is a function for setting other json plugin instead of official json.Unmarshal().
//
// SetUnmarshalFunc 是一個用來設定其他 json 套件來取代官方的 json.Unmarshal() 的函式。
func SetUnmarshalFunc(f func([]byte, interface{}) error) {
	unmarshalFunc = f
}

func callUnmarshal(jsonString []byte, data interface{}) error {
	var err error
	if unmarshalFunc != nil {
		err = unmarshalFunc(jsonString, &data)
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

func isStruct(rv reflect.Value) bool {
	return rv.Kind() == reflect.Struct
}

func isPtrNil(rv reflect.Value) bool {
	return rv.Kind() == reflect.Ptr && rv.IsNil()
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
			return reflect.Value{}, fmt.Errorf("time.Parse error: %w", err)
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
