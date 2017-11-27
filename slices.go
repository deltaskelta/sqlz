package sqlz

import (
	"database/sql/driver"
	"strings"
)

// TODO: test this and handle the case when the slice is nil

// StringSlice is the data type that can be used for marshaling and scanning the values
// into an sql database
type StringSlice []string

// Scan changes the value from a slice of bytes and sets it back as a slice of strings.
// Implements the scanner interface
func (p *StringSlice) Scan(src interface{}) error {
	b := src.([]byte)

	// strings.Split returns a slice with a single empty string when there are no
	// elements to split, so if slice[0] and there is only one element
	slice := strings.Split(string(b), ",")
	if slice[0] == "" && cap(slice) == 1 {
		slice = []string{}
	}

	*p = slice

	if len(*p) == 0 {
		*p = nil
	}

	return nil
}

// Value marshals the slice into a string separated by commas for the database
func (p StringSlice) Value() (driver.Value, error) {
	j := strings.Join([]string(p), ",")
	return j, nil
}
