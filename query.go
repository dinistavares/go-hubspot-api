package hubspot

import (
	"fmt"
	"reflect"
	"strings"
)

type QueryValues map[string]string

// Create Query parameters for accounts routes.
func Query() *QueryValues {
	return &QueryValues{}
}

func (v *QueryValues) getQueryValues() QueryValues {
	return *v
}

func (v QueryValues) encode() string {
	var query string

	count := 0

	for key, value := range v {
		if count > 0 {
			query += "&"
		} else {
			query = "?"
		}

		query += key + "=" + value

		count++
	}

	return query
}

func (v QueryValues) SetArchived(archived bool) {
	v["archived"] = fmt.Sprintf("%t", archived)
}

func (v QueryValues) SetProperties(properties []string) {
	v["properties"] = strings.Join(properties, ",")
}

func (v QueryValues) SetPropertiesWithHistory(properties []string) {
	v["propertiesWithHistory"] = strings.Join(properties, ",")
}

func (v QueryValues) SetCustomPropery(key string, value interface{}) {
	v[key] = fmt.Sprintf("%v", value)
}

func isPointerWithQueryValues(i interface{}) (interface{}, bool) {
	if i == nil {
		return nil, false
	}

	val := reflect.ValueOf(i)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Map {
		return val.Interface(), true
	}

	return nil, false
}

func (v QueryValues) setDeveloperAPIKey(value string) {
	v["hapikey"] = value
}
