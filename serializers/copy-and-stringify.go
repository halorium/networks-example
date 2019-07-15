package serializers

import (
	"fmt"
	"reflect"
)

func CopyAndStringify(rawObject interface{}, newObject interface{}) {
	rawValue := reflect.ValueOf(rawObject)
	newValue := reflect.ValueOf(newObject)

	copyAndStringifyRecursively(rawValue, newValue)
}

func copyAndStringifyRecursively(rawValue, newValue reflect.Value) {
	switch rawValue.Kind() {

	// case Pointer
	case reflect.Ptr:
		rawValueElem := rawValue.Elem()

		if !rawValueElem.IsValid() {
			return
		}

		newValueElem := newValue.Elem()

		if !newValueElem.IsValid() {
			return
		}

		copyAndStringifyRecursively(rawValueElem, newValueElem)

	// case Interface
	case reflect.Interface:
		rawValueElem := rawValue.Elem()
		newValueElem := newValue.Elem()

		if !newValueElem.IsValid() {
			newValueElem = reflect.New(rawValueElem.Type()).Elem()
		}

		copyAndStringifyRecursively(rawValueElem, newValueElem)
		newValue.Set(newValueElem)

	// case Struct
	case reflect.Struct:
		for i := 0; i < rawValue.NumField(); i++ {
			copyAndStringifyRecursively(rawValue.Field(i), newValue.Field(i))
		}

	// case Slice
	case reflect.Slice:
		newValue.Set(reflect.MakeSlice(newValue.Type(), rawValue.Len(), rawValue.Cap()))
		for i := 0; i < rawValue.Len(); i++ {
			copyAndStringifyRecursively(rawValue.Index(i), newValue.Index(i))
		}

	// case Map
	case reflect.Map:
		// Create a new map of newValue Type and translate each value
		newValue.Set(reflect.MakeMap(newValue.Type()))

		// If the type is defferent then we create a new map of rawValue Type
		if rawValue.Kind() != newValue.Kind() {
			newValue.Set(reflect.MakeMap(rawValue.Type()))
		}

		for _, key := range rawValue.MapKeys() {
			rawValueField := rawValue.MapIndex(key)
			newValueField := newValue.MapIndex(key)

			// if not valid then create newValueField from rawValueField
			if !newValueField.IsValid() {
				newValueField = reflect.New(rawValueField.Type()).Elem()
			}

			copyAndStringifyRecursively(rawValueField, newValueField)
			newValue.SetMapIndex(key, newValueField)
		}

		// finished with recursion

	// case Default
	default:
		stringValue := fmt.Sprintf("%v", rawValue.Interface())
		newValue.SetString(stringValue)
	}
}
