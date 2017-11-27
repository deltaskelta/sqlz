package sqlz

import (
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {
	tests := []struct {
		name           string
		data           []byte
		expectedReturn StringSlice
	}{
		{
			name:           "one result",
			data:           []byte("foo"),
			expectedReturn: StringSlice{"foo"},
		},
		{
			name:           "multiple results",
			data:           []byte("foo,bar"),
			expectedReturn: StringSlice{"foo", "bar"},
		},
		{
			name:           "empty string in an invalid way",
			data:           []byte(""),
			expectedReturn: nil,
		},
		{
			name:           "empty string in a valid way",
			data:           []byte("foo,bar,"),
			expectedReturn: StringSlice{"foo", "bar", ""},
		},
		{
			name:           "null data returns nil slice",
			data:           nil,
			expectedReturn: nil,
		},
		{
			name:           "test comma returns two empty strings",
			data:           []byte(","),
			expectedReturn: StringSlice{"", ""},
		},
	}

	for _, tt := range tests {
		var s StringSlice
		err := s.Scan(tt.data)
		if err != nil {
			t.Errorf("%s: failed with err: %v", tt.name, err)
		}

		if !reflect.DeepEqual(s, tt.expectedReturn) {
			t.Errorf("%s: %v != %v", tt.name, s, tt.expectedReturn)
		}
	}
}

func TestValue(t *testing.T) {
	tests := []struct {
		name           string
		data           []string
		expectedReturn string
	}{
		{
			name:           "test multiple values",
			data:           []string{"foo", "bar"},
			expectedReturn: "foo,bar",
		},
		{
			name:           "test single value",
			data:           []string{"foo"},
			expectedReturn: "foo",
		},
		{
			name:           "test empty string",
			data:           []string{""},
			expectedReturn: "",
		},
		{
			name:           "test nil gives empty string",
			data:           nil,
			expectedReturn: "",
		},
	}

	for _, tt := range tests {
		var s = StringSlice(tt.data)

		val, err := s.Value()
		if err != nil {
			t.Errorf("%s: failed with err: %v", tt.name, err)
		}

		if val != tt.expectedReturn {
			t.Errorf("%s: %v != %v", tt.name, val, tt.expectedReturn)
		}
	}
}
