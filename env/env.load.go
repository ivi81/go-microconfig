package env

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// PopulateWithEnv - осуществляет загрузку переменных окружения названия которых указанны ввиде
// значения тэга c названием 'env' поля структуры.
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

		//обработка вложенной структуры
		/*_, ok := typeField.Tag.Lookup("obj")
		if ok && valueField.Kind() == reflect.Struct {
			if err = PopulateWithEnv(prefix, valueField.Addr().Interface()); err != nil {
				return
			}
			continue
		}*/

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

	fieldName := field.Type().Name()

	switch field.Kind() {
	case reflect.String:

		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("error parsing int value for field %s: %v", fieldName, err)
		}
		field.SetInt(intValue)
	case reflect.Float32, reflect.Float64:

		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("error parsing float value for field %s: %v", fieldName, err)
		}
		field.SetFloat(floatValue)
	case reflect.Bool:

		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("error parsing bool value for field %s: %v", fieldName, err)
		}
		field.SetBool(boolValue)
	case reflect.Slice:

		slice := strings.Split(value, strSplitter)
		if len(slice) != 0 {
			s := reflect.ValueOf(slice)
			field.Set(s)
		}
	default:
		return fmt.Errorf("unsupported type for field %s: %s", fieldName, field.Kind())
	}
	return nil
}
