package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/alaa-aqeel/school-managment-system/app/interfaces/observer"
)

func ExcuteMethod[T any](observer observer.Observer[T], name string, model T) error {
	observerValue := reflect.ValueOf(observer)
	method := observerValue.MethodByName(name)
	if !method.IsValid() {

		return errors.New("[Error]: Method does not exist")
	}
	modelValue := reflect.ValueOf(model)
	if method.Type().NumIn() != 1 {
		return errors.New("[Error]: Method does not accept the correct number of arguments")
	}
	if method.Type().In(0).Kind() == reflect.Ptr && modelValue.Kind() != reflect.Ptr {
		modelValue = reflect.New(modelValue.Type()).Elem()
		modelValue.Set(reflect.ValueOf(model))
	}
	method.Call([]reflect.Value{modelValue})
	return nil
}

func ExceuteObservers[T any](observers []observer.Observer[T], name string, model T) error {
	for _, observer := range observers {
		if err := ExcuteMethod(observer, name, model); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
