package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// change instance=["a", "b"]
// to instance.1="a" instance.2="b"
func FlattenFn(fieldName string, field reflect.Value, values *url.Values) {
	l := field.Len()
	if l > 0 {
		for i := 0; i < l; i++ {
			str := field.Index(i).String()
			values.Set(fieldName+"."+strconv.Itoa(i+1), str)
		}
	}
}

func FlattenMultiDimensionalFn(fieldName string, field reflect.Value, values *url.Values) {
	l := field.Len()
	if l > 0 {
		for i := 0; i < l; i++ {
			subArray := field.Index(i)
			J := subArray.Len()
			for j := 0; j < J; j++ {
				str := subArray.Index(j).String()
				if j == 0 {
					values.Set(fieldName+"."+strconv.Itoa(i)+"."+"name", str)
				}
				if j == 1 {
					values.Set(fieldName+"."+strconv.Itoa(i)+"."+"value", str)
				}
			}
		}
	}
}

//ConvertToParamValues converts the struct to url.Values
func ConvertToparamValues(ifc interface{}, uppercase bool) url.Values {
	values := url.Values{}
	SetParamValues(ifc, &values, uppercase)
	return values
}

//SetParamValues sets the struct to existing url.Values following ECS encoding rules
func SetParamValues(ifc interface{}, values *url.Values, uppercase bool) {
	setParamValues(ifc, values, "", uppercase)
}

func makeFirstLowerCase(s string, uppercase bool) string {

	if uppercase {
		return s
	} else {

		if len(s) < 2 {
			return strings.ToLower(s)
		}

		bts := []byte(s)

		lc := bytes.ToLower([]byte{bts[0]})
		rest := bts[1:]

		return string(bytes.Join([][]byte{lc, rest}, nil))
	}

}

func setParamValues(i interface{}, values *url.Values, prefix string, uppercase bool) {
	// add to support url.Values

	mapValues, ok := i.(url.Values)
	if ok {
		for k := range mapValues {
			values.Set(k, mapValues.Get(k))
		}
		return
	}

	elem := reflect.ValueOf(i)

	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	elemType := elem.Type()
	for i := 0; i < elem.NumField(); i++ {

		fieldName := makeFirstLowerCase(elemType.Field(i).Name, uppercase)

		anonymous := elemType.Field(i).Anonymous
		field := elem.Field(i)

		//fmt.Println(values)

		// TODO Use Tag for validation
		// tag := typ.Field(i).Tag.Get("tagname")
		kind := field.Kind()

		if (kind == reflect.Ptr || kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map || kind == reflect.Chan) && field.IsNil() {
			continue
		}
		if kind == reflect.Ptr {
			field = field.Elem()
			kind = field.Kind()
		}

		var value string
		//switch field.Interface().(type) {
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i := field.Int()
			if i != 0 {
				value = strconv.FormatInt(i, 10)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			i := field.Uint()
			if i != 0 {
				value = strconv.FormatUint(i, 10)
			}
		case reflect.Float32:
			value = strconv.FormatFloat(field.Float(), 'f', 4, 32)
		case reflect.Float64:
			value = strconv.FormatFloat(field.Float(), 'f', 4, 64)
		case reflect.Bool:
			value = strconv.FormatBool(field.Bool())
		case reflect.String:
			value = field.String()
		case reflect.Map:
			ifc := field.Interface()
			m := ifc.(map[string]string)
			if m != nil {
				j := 0
				for k, v := range m {
					j++
					keyName := fmt.Sprintf("%s.%d.Key", fieldName, j)
					values.Set(keyName, k)
					valueName := fmt.Sprintf("%s.%d.Value", fieldName, j)
					values.Set(valueName, v)
				}
			}
		case reflect.Slice:
			if field.Type().Name() == "FlattenArray" {
				FlattenFn(fieldName, field, values)
			} else if field.Type().Name() == "FlattenMultiDimensionalArray" {
				FlattenMultiDimensionalFn(fieldName, field, values)
			} else {
				switch field.Type().Elem().Kind() {
				case reflect.Uint8:
					value = string(field.Bytes())
				case reflect.String:
					l := field.Len()
					if l > 0 {
						strArray := make([]string, l)
						for i := 0; i < l; i++ {
							strArray[i] = field.Index(i).String()
						}
						bytes, err := json.Marshal(strArray)
						if err == nil {
							value = string(bytes)
						} else {
							log.Printf("Failed to convert JSON: %v", err)
						}
					}
				default:
					l := field.Len()
					for j := 0; j < l; j++ {
						prefixName := fmt.Sprintf("%s.%d.", fieldName, (j + 1))
						ifc := field.Index(j).Interface()
						//log.Printf("%s : %v", prefixName, ifc)
						if ifc != nil {
							setParamValues(ifc, values, prefixName, uppercase)
						}
					}
					continue
				}
			}

		default:
			switch field.Interface().(type) {
			case time.Time:
				t := field.Interface().(time.Time)
				value = t.String()
			default:

				ifc := field.Interface()
				if ifc != nil {
					if anonymous {
						SetParamValues(ifc, values, uppercase)
					} else {
						prefixName := fieldName + "."
						setParamValues(ifc, values, prefixName, uppercase)
					}
					continue
				}
			}
		}
		if value != "" {
			name := elemType.Field(i).Tag.Get("ArgName")
			if name == "" {
				name = fieldName
			}
			if prefix != "" {
				name = prefix + name
			}
			// NOTE: here we will change name to underline style when the type is UnderlineString
			if field.Type().Name() == "UnderlineString" {
				name = Underline2Dot(name)
			}
			values.Set(name, value)
		}
	}
}

func CreateRandomString() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

func Underline2Dot(name string) string {
	return strings.Replace(name, "_", ".", -1)
}
