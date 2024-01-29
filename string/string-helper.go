package stringHelper

import "reflect"

// Function to check string value is valid
func IsValidString(value string) bool {
	if len(value) == 0 {
		return false
	}

	return true
}

// Function to check if a field is empty based on its type
func IsEmpty(value interface{}) bool {
	_, ok := value.(reflect.Value)
	var v reflect.Value
	if ok {
		v = value.(reflect.Value)
	} else {
		v = reflect.ValueOf(value)
	}

	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Struct:
		// Check if struct is empty by comparing each field with the zero value for its type
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			zeroValue := reflect.Zero(field.Type())
			if !reflect.DeepEqual(field.Interface(), zeroValue.Interface()) {
				return false
			}
		}
		return true
	}

	return false
}

// Function to check if a field is empty based on its type
func IsValidFields(fields []string, values []interface{}) (bool, string) {
	for index, value := range values {
		if IsEmpty(value) {
			return false, fields[index]
		}
	}
	return true, ""
}
