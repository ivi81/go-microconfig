// env.load.go - содержит функции осуществлющие загрузку переменных окружения в структуру данных
package env

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ivi81/enummethods/enumerator"
)

// Enumerator интерфейсный для работы с нестроковыми константами
type Enumerator interface {
	enumerator.Unstringer
	enumerator.Validator
}

// PopulateWithEnv - осуществляет загрузку переменных окружения названия которых указанны ввиде
// значения тэга 'env' для поля структуры.
//
//	параметры:
//	prefix - префикс названия переменной
//	s - заполняемая структура данных
func PopulateWithEnv(prefix string, s any) (err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	val := reflect.ValueOf(s).Elem()

	for i := 0; i < val.NumField(); i++ {

		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		if tag, ok := typeField.Tag.Lookup("env"); ok {
			fields := strings.Split(tag, ",")
			var envKey string

			if len(fields) > 0 && fields[0] != "_" {
				envKey = JoinStr(prefix, fields[0])
			} else {
				envKey = prefix
			}

			if valueField.Kind() == reflect.Struct {

				f := valueField.Addr().Interface()
				if err = PopulateWithEnv(envKey, f); err != nil {
					return
				}
				continue
			}

			if strVal, ok := LookupEnv(envKey); ok {
				if err = assignValue(&valueField, strVal); err != nil {
					return fmt.Errorf("Error assigning value, '%s'", err)
				}
			}
		}
	}

	return
}

// assignValue - приведение строкового значения к типу данных поля заполняемой структуры
func assignValue(field *reflect.Value, value string) error {
	if !field.CanSet() {
		return fmt.Errorf("field is not settable")
	}
	fieldType := field.Type()
	fieldTypeName := fieldType.Name()

	if field.CanAddr() {
		addr := field.Addr()
		if addr.IsValid() {
			// Проверяем интерфейс Enumerator
			if enumerator, ok := addr.Interface().(Enumerator); ok {
				if enumerator.SetValue(value) {
					if !enumerator.IsValid() {
						return fmt.Errorf("not valid value %s for type %s", value, fieldTypeName)
					}
					return nil
				}
				return fmt.Errorf("SetValue failed for value %s and type %s", value, fieldTypeName)
			}
		}
	}

	switch field.Kind() {
	case reflect.String:

		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		if fieldType == reflect.TypeOf(time.Duration(0)) { //Тут обрабатываем если значение поля должно быть временным интервалом
			timeValue, err := time.ParseDuration(value)
			if err != nil {
				return fmt.Errorf("error parsing time value for field %s: %v", fieldTypeName, err)
			}
			field.SetInt(int64(timeValue))

			/*if fieldTypeName == "Duration" { //Тут обрабатываем если значение поля должно быть временным интервалом
			timeValue, err := time.ParseDuration(value)
			if err != nil {
				return fmt.Errorf("error parsing time value for field %s: %v", fieldTypeName, err)
			}
			field.SetInt(int64(timeValue))*/

		} else {

			intValue, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("error parsing int value for field %s: %v", fieldTypeName, err)
			}
			field.SetInt(intValue)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return fmt.Errorf("error parsing uint value: %w", err)
		}
		field.SetUint(uintValue)
	case reflect.Float32, reflect.Float64:

		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("error parsing float value for field %s: %v", fieldTypeName, err)
		}
		field.SetFloat(floatValue)
	case reflect.Bool:

		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("error parsing bool value for field %s: %v", fieldTypeName, err)
		}
		field.SetBool(boolValue)
	case reflect.Slice:

		slice := strings.Split(value, strSplitter)
		if len(slice) != 0 {
			s := reflect.ValueOf(slice)
			field.Set(s)
		}
	default:
		return fmt.Errorf("unsupported type for field %s: %s", fieldTypeName, field.Kind())
	}
	return nil
}
