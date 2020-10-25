package api

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type InputTransformError struct {
	code  int
	field string
	desc  string
}

const (
	ErrInputTransformCodeMissing = iota
	ErrInputTransformCodeTypeMismatch
)

func newInputTransformError(code int, field, desc string) *InputTransformError {
	return &InputTransformError{
		code:  code,
		field: field,
		desc:  desc,
	}
}

func (err *InputTransformError) Code() int {
	return err.code
}

func (err *InputTransformError) Error() string {
	return fmt.Sprintf("Input Parameter[%s] is invalid since %s", err.field, err.desc)
}

var (
	ErrInputTransformExpectStruct = errors.New("target must be a struct")
)

// tagOptions is the string following a comma in a struct field's "json"
// tag, or the empty string. It does not include the leading comma.
type tagOptions string

// parseTag splits a struct field's json tag into its name and
// comma-separated options.
func parseTag(tag string) (string, tagOptions) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tagOptions(tag[idx+1:])
	}
	return tag, tagOptions("")
}

// Contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var next string
		i := strings.Index(s, ",")
		if i >= 0 {
			s, next = s[:i], s[i+1:]
		}
		if s == optionName {
			return true
		}
		s = next
	}
	return false
}

func transformLog(message map[string]interface{}, action string) {
	//log.ObjectDebug(message, "[transform-input] "+action)
	//fmt.Println(time.Now().Format(SIMPLE_TIME_FORMAT), "[transform-input] "+action, message)
}

func TransformInput(input map[string]interface{}, ty reflect.Type) (interface{}, error) {
	transformLog(map[string]interface{}{
		"input": input,
	}, "begin")
	if ty.Kind() != reflect.Struct {
		return nil, errors.New("transform only accepts struct")
	}
	target := reflect.New(ty).Elem()
	fieldCount := ty.NumField()

	for i := 0; i < fieldCount; i++ {
		field := target.Field(i)
		fieldType := ty.Field(i)
		transformLog(map[string]interface{}{
			"fieldType": fieldType,
		}, "fieldType")

		name, optional := parseJsonAnotation(fieldType)
		transformLog(map[string]interface{}{
			"fieldName": name,
		}, "parse json anotation")

		sName := strings.Split(name, ",")
		transformLog(map[string]interface{}{
			"sName": sName,
		}, "split string sName")

		iName := sName[0]
		if v, ok := input[iName]; ok {
			err := transformSingleField(v, &field)
			if err != nil {
				return nil, err
			}
		} else {
			if optional {
				continue
			}
			return nil, newInputTransformError(ErrInputTransformCodeMissing, name, "Missing in Input")
		}
	}
	transformLog(map[string]interface{}{
		"target": target,
	}, "end")
	return target.Interface(), nil
}

func parseJsonAnotation(field reflect.StructField) (name string, optional bool) {
	optional = false
	j := field.Tag.Get("json")
	_, opts := parseTag(j)
	if opts.Contains("omitempty") {
		optional = true
	}
	if j == "" {
		name = field.Name
	} else {
		name = j
	}
	return
}

func transformSingleField(v interface{}, val *reflect.Value) error {
	transformLog(map[string]interface{}{
		"kind": val.Kind(),
	}, "type kind")
	switch val.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if n, err := transformToUint(v); err == nil {
			val.SetUint(n)
			return nil
		} else {
			return err
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if n, err := transformToInt(v); err == nil {
			val.SetInt(n)
			return nil
		} else {
			return err
		}
	case reflect.Float32, reflect.Float64:
		if f, err := transformToFloat(v); err == nil {
			val.SetFloat(f)
			return nil
		} else {
			return err
		}
	case reflect.Bool:
		if b, err := transformToBool(v); err == nil {
			val.SetBool(b)
			return nil
		} else {
			return err
		}
	case reflect.String:
		if vIn, ok := v.(string); ok {
			val.SetString(vIn)
			return nil
		} else {
			return errors.New("input with required type not match, please check")
		}
	case reflect.Slice:
		if vInArr, ok := v.([]interface{}); ok {
			str := make([]string, len(vInArr))
			for i, vIn := range vInArr {
				//log.DebugSimple("in []string", vIn.(string))
				str[i] = vIn.(string)
			}
			val.Set(reflect.ValueOf(str))
			return nil
		} else {
			return errors.New("parse slice error, please check")
		}
	default:
		return errors.New(fmt.Sprintf("can't understand this type:%s", val.Kind()))
	}
}

func transformToUint(value interface{}) (uint64, error) {
	valueType := reflect.TypeOf(value)
	transformLog(map[string]interface{}{
		"kind": valueType.Kind(),
	}, "transform to uint")
	switch valueType.Kind() {
	case reflect.Float64:
		return uint64(value.(float64)), nil
	case reflect.String:
		if t, err := strconv.ParseUint(value.(string), 10, 64); err == nil {
			return uint64(t), nil
		} else {
			return uint64(0), err
		}
	default:
		return uint64(0), errors.New("input type error")
	}
}

func transformToBool(value interface{}) (bool, error) {
	valueType := reflect.TypeOf(value)
	transformLog(map[string]interface{}{
		"kind": valueType.Kind(),
	}, "transform to bool")
	switch valueType.Kind() {
	case reflect.Bool:
		return value.(bool), nil
	case reflect.String:
		if b, err := strconv.ParseBool(value.(string)); err == nil {
			return b, nil
		} else {
			return false, err
		}
	default:
		return false, errors.New("input type error")

	}
}

func transformToInt(value interface{}) (int64, error) {
	valueType := reflect.TypeOf(value)
	transformLog(map[string]interface{}{
		"kind": valueType.Kind(),
	}, "transform to int")
	switch valueType.Kind() {
	case reflect.Float64:
		return int64(value.(float64)), nil
	case reflect.String:
		if t, err := strconv.ParseInt(value.(string), 10, 64); err == nil {
			return int64(t), nil
		} else {
			return int64(0), err
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int64(value.(int64)), nil
	default:
		return int64(0), errors.New("input type error")
	}
}

func transformToFloat(value interface{}) (float64, error) {
	valueType := reflect.TypeOf(value)
	transformLog(map[string]interface{}{
		"kind": valueType.Kind(),
	}, "transform to float")
	switch valueType.Kind() {
	case reflect.Float64:
		return value.(float64), nil
	case reflect.String:
		if t, err := strconv.ParseFloat(value.(string), 64); err == nil {
			return float64(t), nil
		} else {
			return 0.0, nil
		}
	default:
		return 0.0, errors.New("input type error")
	}
}
