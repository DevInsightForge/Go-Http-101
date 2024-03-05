package utility

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func ParseQueryParams(queryParams url.Values, options interface{}) {
	val := reflect.Indirect(reflect.ValueOf(options))
	if val.Kind() != reflect.Struct {
		return
	}

	normalizedParams := normalizeQueryParams(queryParams)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.CanSet() {
			continue
		}

		fieldName := strings.ToLower(val.Type().Field(i).Name)
		values, exists := normalizedParams[fieldName]
		if !exists || len(values) == 0 {
			continue
		}

		setFieldValueFromQuery(field, values[0])
	}
}

func normalizeQueryParams(queryParams url.Values) map[string][]string {
	normalized := make(map[string][]string)
	for key, values := range queryParams {
		normalized[strings.ToLower(key)] = values
	}
	return normalized
}

func setFieldValueFromQuery(field reflect.Value, value string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if parsed, err := strconv.ParseInt(value, 10, 64); err == nil {
			field.SetInt(parsed)
		}
	case reflect.Bool:
		if parsed, err := strconv.ParseBool(value); err == nil {
			field.SetBool(parsed)
		}
	case reflect.Float32, reflect.Float64:
		if parsed, err := strconv.ParseFloat(value, 64); err == nil {
			field.SetFloat(parsed)
		}
	}
}
