package sqlz

import (
	"testing"
)

func TestUpdateSQL(t *testing.T) {
	tests := []struct {
		name        string
		struc       interface{}
		table       string
		expected    string
		expectedErr error
	}{
		{
			name: "test successful partial update sql",
			struc: struct {
				ID     int64 `db:"id"`
				Field1 int   `db:"field_1"`
			}{
				ID:     1,
				Field1: 1,
			},
			table:       "test",
			expected:    "UPDATE test SET field_1 = :field_1 WHERE id = :id",
			expectedErr: nil,
		},
		{
			name: "test with multiple fields",
			struc: struct {
				ID     int64  `db:"id"`
				Field1 int    `db:"field_1"`
				Field2 string `db:"s"`
			}{
				ID:     1,
				Field1: 1,
				Field2: "one",
			},
			table:       "test",
			expected:    "UPDATE test SET field_1 = :field_1,s = :s WHERE id = :id",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		sql, err := UpdateSQL(tt.struc, tt.table)
		if err != tt.expectedErr {
			t.Errorf("%s: expected err: %v got err: %v", tt.name, tt.expectedErr, err)
		}
		if err != nil {
			continue // got expected err and err != nil so nothing left to test
		}

		if sql != tt.expected {
			t.Errorf("%s:\n%s\nDOES NOT EQUAL\n%s\n", tt.name, sql, tt.expected)
		}

	}
}

func TestUpdatedFields(t *testing.T) {
	tests := []struct {
		name        string
		struc       interface{}
		table       string
		expected    string
		expectedErr error
	}{
		{
			name: "test successful partial update",
			struc: struct {
				ID     int64
				Field1 int `db:"field_1"`
			}{
				ID:     1,
				Field1: 1,
			},
			expectedErr: ErrNoStructTag,
		},
		{
			name: "test zero id gives error",
			struc: struct {
				ID     int64 `db:"id"`
				Field1 int   `db:"field_1"`
			}{
				Field1: 1,
			},
			table:       "test",
			expectedErr: ErrZeroID,
		},
		{
			name: "test with multiple fields",
			struc: struct {
				ID     int64    `db:"id"`
				Field1 int      `db:"a"`
				Field2 string   `db:"b"`
				Field3 []string `db:"c"`
			}{
				ID:     1,
				Field1: 1,
				Field2: "one",
			},
			table:       "test",
			expected:    "a = :a,b = :b",
			expectedErr: nil,
		},
		{
			name: "test with multiple fields",
			struc: struct {
				ID     int64    `db:"id"`
				Field1 int      `db:"a"`
				Field2 string   `db:"b"`
				Field3 []string `db:"c"`
			}{
				ID:     1,
				Field1: 1,
				Field2: "one",
				Field3: []string{"notempty"},
			},
			table:       "test",
			expected:    "a = :a,b = :b,c = :c",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		sql, err := UpdatedFields(tt.struc)
		if err != tt.expectedErr {
			t.Errorf("%s: expected err: %v got err: %v", tt.name, tt.expectedErr, err)
		}
		if err != nil {
			continue // got expected err and err != nil so nothing left to test
		}

		if sql != tt.expected {
			t.Errorf("%s:\n%s\nDOES NOT EQUAL\n%s\n", tt.name, sql, tt.expected)
		}

	}
}
