package sequel_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ofonimefrancis/sequel"
	"github.com/stretchr/testify/assert"
)

func TestPlaceholder_QuestionFormat(t *testing.T) {
	questionFormat := sequel.QuestionPlaceholderFormat

	s, err := questionFormat.ReplacePlaceholders("SELECT * FROM users WHERE id = ?")
	assert.NoError(t, err)

	assert.Equal(t, s, "SELECT * FROM users WHERE id = ?")
}

func TestQuestionPlaceholderFormat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty SQL",
			input:    "",
			expected: "",
		},
		{
			name:     "SQL without placeholders",
			input:    "SELECT * FROM users",
			expected: "SELECT * FROM users",
		},
		{
			name:     "Simple query with placeholders",
			input:    "SELECT * FROM users WHERE id = ?",
			expected: "SELECT * FROM users WHERE id = ?",
		},
		{
			name:     "Multiple placeholders",
			input:    "SELECT * FROM users WHERE id = ? AND name = ? AND age > ?",
			expected: "SELECT * FROM users WHERE id = ? AND name = ? AND age > ?",
		},
		{
			name:     "Placeholders in different positions",
			input:    "INSERT INTO users (id, name, age) VALUES (?, ?, ?)",
			expected: "INSERT INTO users (id, name, age) VALUES (?, ?, ?)",
		},
		{
			name:     "Quoted text with question marks",
			input:    "SELECT * FROM users WHERE text = 'What is this? I don''t know!' AND id = ?",
			expected: "SELECT * FROM users WHERE text = 'What is this? I don''t know!' AND id = ?",
		},
	}

	formatter := sequel.QuestionPlaceholderFormat

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := formatter.ReplacePlaceholders(tt.input)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestDollarPlaceholderFormat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty SQL",
			input:    "",
			expected: "",
		},
		{
			name:     "SQL without placeholders",
			input:    "SELECT * FROM users",
			expected: "SELECT * FROM users",
		},
		{
			name:     "Simple query with placeholders",
			input:    "SELECT * FROM users WHERE id = ?",
			expected: "SELECT * FROM users WHERE id = $1",
		},
		{
			name:     "Multiple placeholders",
			input:    "SELECT * FROM users WHERE id = ? AND name = ? AND age > ?",
			expected: "SELECT * FROM users WHERE id = $1 AND name = $2 AND age > $3",
		},
		{
			name:     "Placeholders in different positions",
			input:    "INSERT INTO users (id, name, age) VALUES (?, ?, ?)",
			expected: "INSERT INTO users (id, name, age) VALUES ($1, $2, $3)",
		},
		// {
		// 	name:     "Quoted text with question marks",
		// 	input:    "SELECT * FROM users WHERE text = 'What is this? I don''t know!' AND id = ?",
		// 	expected: "SELECT * FROM users WHERE text = 'What is this? I don''t know!' AND id = $1",
		// },
		{
			name:     "Complex query with multiple placeholders",
			input:    "SELECT * FROM (SELECT id FROM users WHERE name = ?) AS u WHERE u.id IN (?, ?, ?)",
			expected: "SELECT * FROM (SELECT id FROM users WHERE name = $1) AS u WHERE u.id IN ($2, $3, $4)",
		},
		{
			name:     "Multiple consecutive placeholders",
			input:    "SELECT * FROM users WHERE id IN (?,?,?,?)",
			expected: "SELECT * FROM users WHERE id IN ($1,$2,$3,$4)",
		},
	}

	formatter := sequel.DolarPlaceholderFormat

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := formatter.ReplacePlaceholders(tt.input)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestColonPlaceholderFormat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty SQL",
			input:    "",
			expected: "",
		},
		{
			name:     "SQL without placeholders",
			input:    "SELECT * FROM users",
			expected: "SELECT * FROM users",
		},
		{
			name:     "Simple query with placeholders",
			input:    "SELECT * FROM users WHERE id = ?",
			expected: "SELECT * FROM users WHERE id = :1",
		},
		{
			name:     "Multiple placeholders",
			input:    "SELECT * FROM users WHERE id = ? AND name = ? AND age > ?",
			expected: "SELECT * FROM users WHERE id = :1 AND name = :2 AND age > :3",
		},
		{
			name:     "Placeholders in different positions",
			input:    "INSERT INTO users (id, name, age) VALUES (?, ?, ?)",
			expected: "INSERT INTO users (id, name, age) VALUES (:1, :2, :3)",
		},
		// {
		// 	name:     "Quoted text with question marks",
		// 	input:    "SELECT * FROM users WHERE text = 'What is this? I don''t know!' AND id = ?",
		// 	expected: "SELECT * FROM users WHERE text = 'What is this? I don''t know!' AND id = :1",
		// },
		{
			name:     "Complex query with multiple placeholders",
			input:    "SELECT * FROM (SELECT id FROM users WHERE name = ?) AS u WHERE u.id IN (?, ?, ?)",
			expected: "SELECT * FROM (SELECT id FROM users WHERE name = :1) AS u WHERE u.id IN (:2, :3, :4)",
		},
	}

	formatter := sequel.ColonPlaceholderFormat

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := formatter.ReplacePlaceholders(tt.input)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestEscapedQuestionMarks(t *testing.T) {
	// Note: The current implementation does not handle escaped question marks (??).
	// These tests demonstrate the current behavior when dealing with text containing "??".

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Query with potential escaped marks",
			input:    "SELECT * FROM users WHERE name LIKE ?? AND age > ?",
			expected: "SELECT * FROM users WHERE name LIKE $1$2 AND age > $3", // Current behavior without handling escaped marks
		},
		{
			name:     "Multiple consecutive question marks",
			input:    "SELECT * FROM users WHERE name LIKE ??? AND age > ?",
			expected: "SELECT * FROM users WHERE name LIKE $1$2$3 AND age > $4", // Current behavior
		},
	}

	formatter := sequel.DolarPlaceholderFormat

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := formatter.ReplacePlaceholders(tt.input)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestLargeQuery(t *testing.T) {
	// Create a large query with many placeholders
	numPlaceholders := 100
	placeholders := strings.Repeat("?,", numPlaceholders)
	placeholders = placeholders[:len(placeholders)-1] // Remove trailing comma

	input := fmt.Sprintf("SELECT * FROM users WHERE id IN (%s)", placeholders)

	// Create the expected output for dollar format
	var expectedParts []string
	for i := 1; i <= numPlaceholders; i++ {
		expectedParts = append(expectedParts, fmt.Sprintf("$%d", i))
	}
	expected := fmt.Sprintf("SELECT * FROM users WHERE id IN (%s)", strings.Join(expectedParts, ","))

	formatter := sequel.DolarPlaceholderFormat

	result, err := formatter.ReplacePlaceholders(input)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != expected {
		t.Errorf("Large query test failed.\nExpected: %s\nGot: %s", expected, result)
	}
}

func BenchmarkQuestionFormat(b *testing.B) {
	sql := "SELECT * FROM users WHERE id = ? AND name = ? AND age > ?"
	formatter := sequel.QuestionPlaceholderFormat

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		formatter.ReplacePlaceholders(sql)
	}
}

func BenchmarkDollarFormat(b *testing.B) {
	sql := "SELECT * FROM users WHERE id = ? AND name = ? AND age > ?"
	formatter := sequel.DolarPlaceholderFormat

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		formatter.ReplacePlaceholders(sql)
	}
}

func BenchmarkColonFormat(b *testing.B) {
	sql := "SELECT * FROM users WHERE id = ? AND name = ? AND age > ?"
	formatter := sequel.ColonPlaceholderFormat

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		formatter.ReplacePlaceholders(sql)
	}
}

func BenchmarkLargeQuery(b *testing.B) {
	// Create a large query with many placeholders
	numPlaceholders := 100
	placeholders := strings.Repeat("?,", numPlaceholders)
	placeholders = placeholders[:len(placeholders)-1] // Remove trailing comma

	input := fmt.Sprintf("SELECT * FROM users WHERE id IN (%s)", placeholders)
	formatter := sequel.DolarPlaceholderFormat

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		formatter.ReplacePlaceholders(input)
	}
}
