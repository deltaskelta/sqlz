// sqlz provides a few helper functions for writing sql in conjunction with
// github.com/jmoiron/sqlx/
package sqlz

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// error values
var (
	ErrZeroID      = errors.New("structx: id cannot be 0")
	ErrNoStructTag = errors.New("structx: every field needs a `db:\"\"` tag")
)

// UpdateSQL generates sql to update a table based on fields of a struct. If the field is
// a zero value, then it will be left out of the sql query. The generated sql will be a
// fully formed string and it assumes there is an id field and explicit `db:""` tags.
// example output: UPDATE table SET field_1 = :field_1,field_2 = :field_2 WHERE id = :id
func UpdateSQL(s interface{}, table string) (string, error) {
	updFields, err := UpdatedFields(s)
	if err != nil {
		return "", err
	}

	sql := fmt.Sprintf(`UPDATE %s SET %s WHERE id = :id`, table, updFields)

	return sql, nil
}

// UpdatedFields returns partial sql (in the format used by github.com/jmoiron/sqlx)
// it will return a string for every non zero value field of a struct
// example output: field_1 = :field_1,field_2 = :field_2
func UpdatedFields(s interface{}) (string, error) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	f := []string{}

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		valueType := value.Type()

		tagVal := t.Field(i).Tag.Get("db")
		if tagVal == "" {
			return "", ErrNoStructTag
		}

		if !valueType.Comparable() {
			if !value.IsNil() && value.Len() != 0 {
				f = append(f, fmt.Sprintf("%s = :%s", tagVal, tagVal))
			}
			continue
		}

		valueInterface := v.Field(i).Interface()
		zeroValue := reflect.Zero(v.Field(i).Type()).Interface()

		if tagVal == "id" {
			if valueInterface == zeroValue {
				return "", ErrZeroID
			}
			continue
		}

		if valueInterface != zeroValue {
			f = append(f, fmt.Sprintf("%s = :%s", tagVal, tagVal))
		}
	}

	return strings.Join(f, ","), nil
}
