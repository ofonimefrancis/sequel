package sequel

import (
	"reflect"
	"testing"
)

func TestTableName(t *testing.T) {
	tests := []struct {
		name     string
		table    Table
		expected string
	}{
		{
			name:     "regular table name",
			table:    Table{name: "users"},
			expected: "users",
		},
		{
			name:     "empty table name",
			table:    Table{name: ""},
			expected: "",
		},
		{
			name:     "table name with special characters",
			table:    Table{name: "user_data"},
			expected: "user_data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.table.Name()
			if result != tt.expected {
				t.Errorf("Name() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestNewTable(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Table
	}{
		{
			name:     "regular table name",
			input:    "products",
			expected: Table{name: "products"},
		},
		{
			name:     "empty table name",
			input:    "",
			expected: Table{name: ""},
		},
		{
			name:     "table name with special characters",
			input:    "order_items",
			expected: Table{name: "order_items"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := newTable(tt.input)
			if result.name != tt.expected.name {
				t.Errorf("newTable(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTableToSql(t *testing.T) {
	tests := []struct {
		name          string
		table         Table
		expectedSQL   string
		expectedArgs  []interface{}
		expectedError error
	}{
		{
			name:          "regular table",
			table:         Table{name: "customers"},
			expectedSQL:   "customers",
			expectedArgs:  nil,
			expectedError: nil,
		},
		{
			name:          "empty table name",
			table:         Table{name: ""},
			expectedSQL:   "",
			expectedArgs:  nil,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql, args, err := tt.table.ToSql()

			if sql != tt.expectedSQL {
				t.Errorf("ToSql() sql = %v, want %v", sql, tt.expectedSQL)
			}

			if !reflect.DeepEqual(args, tt.expectedArgs) {
				t.Errorf("ToSql() args = %v, want %v", args, tt.expectedArgs)
			}

			if err != tt.expectedError {
				t.Errorf("ToSql() err = %v, want %v", err, tt.expectedError)
			}
		})
	}
}
