// env.load.go - содержит функции осуществлющие загрузку переменных окружения в структуру данных
package env

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"sourcecraft.dev/ivi-ippolitov/enummethods/enumerator"
)

// Enumerator интерфейсный для работы с нестроковыми константами
type Enumerator interface {
	enumerator.Enumerator
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
			err = fmt.Errorf("panic in PopulateWithEnv: %v", r)
		}
	}()

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return fmt.Errorf("expected non-nil pointer to struct")
	}

	val = val.Elem()

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

			// Обработка вложенных структур
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

	return nil
}

// assignValue - приведение строкового значения к типу данных поля заполняемой структуры
func assignValue(field *reflect.Value, value string) error {
	if !field.CanSet() {
		return fmt.Errorf("field is not settable")
	}

	fieldType := field.Type()
	// Универсальная проверка на Enumerator (для одиночных значений и слайсов)
	if isEnumeratorType(fieldType) {
		return setEnumeratorValue(field, value)
	}

	switch field.Kind() {
	case reflect.String:

		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		if fieldType == reflect.TypeOf(time.Duration(0)) { // Тут обрабатываем если значение поля должно быть временным интервалом
			timeValue, err := time.ParseDuration(value)
			if err != nil {
				return fmt.Errorf("error parsing duration: %w", err)
			}
			field.SetInt(int64(timeValue))

		} else {

			intValue, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("error parsing int: %w", err)
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
			return fmt.Errorf("error parsing float: %w", err)
		}
		field.SetFloat(floatValue)
	case reflect.Bool:

		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("error parsing bool %w", err)
		}
		field.SetBool(boolValue)
	case reflect.Slice:
		elemType := fieldType.Elem()
		if isEnumeratorType(elemType) {
			return setEnumeratorSlice(field, value, elemType)
		}

		// Стандартная обработка для []string
		slice := strings.Split(value, strSplitter)
		if len(slice) != 0 {
			s := reflect.ValueOf(slice)
			field.Set(s)
		}
		return nil
	default:

		return fmt.Errorf("unsupported type: %s", fieldType)
	}
	return nil
}

// Проверяет, реализует ли тип интерфейс Enumerator
func isEnumeratorType(t reflect.Type) bool {
	enumeratorInterfaceType := reflect.TypeOf((*Enumerator)(nil)).Elem()
	if t.Implements(enumeratorInterfaceType) {
		return true
	}
	// Если это не указатель, но методы с pointer receiver - создаем указатель
	if t.Kind() != reflect.Pointer {
		ptrType := reflect.PointerTo(t)
		if ptrType.Implements(enumeratorInterfaceType) {
			return true
		}
	}

	return false
}

// Универсальная установка значения для Enumerator
func setEnumeratorValue(field *reflect.Value, value string) error {
	var enumerator Enumerator

	if field.CanAddr() {
		// Для существующего значения
		addr := field.Addr()
		if addr.IsValid() {
			if e, ok := addr.Interface().(Enumerator); ok {
				enumerator = e
			}
		}
	} else {
		// Для создания нового значения
		newVal := reflect.New(field.Type()).Elem()
		if e, ok := newVal.Addr().Interface().(Enumerator); ok {
			enumerator = e
			field.Set(newVal)
		}
	}

	if enumerator != nil {
		if enumerator.SetValue(value) {
			if !enumerator.IsValid() {
				return fmt.Errorf("not valid value %s for type %s", value, field.Type().Name())
			}
			return nil
		}
		return fmt.Errorf("SetValue failed for value %s and type %s", value, field.Type().Name())
	}

	return fmt.Errorf("type does not implement Enumerator")
}

// Функция для слайсов Enumerator'ов (аналогичная первому способу)
func setEnumeratorSlice(field *reflect.Value, value string, elemType reflect.Type) error {
	values := strings.Split(value, strSplitter)
	if len(values) == 0 {
		return nil
	}

	sliceValue := reflect.MakeSlice(field.Type(), len(values), len(values))

	for i, val := range values {
		elem := reflect.New(elemType).Elem()
		elemPtr := elem.Addr()

		if enumerator, ok := elemPtr.Interface().(Enumerator); ok {
			if !enumerator.SetValue(strings.TrimSpace(val)) {
				return fmt.Errorf("SetValue failed for value '%s' at position %d", val, i)
			}
			if !enumerator.IsValid() {
				return fmt.Errorf("not valid value '%s' at position %d for type %s", val, i, elemType.Name())
			}
			sliceValue.Index(i).Set(elem)
		} else {
			return fmt.Errorf("element type does not implement Enumerator")
		}
	}

	field.Set(sliceValue)
	return nil
}
