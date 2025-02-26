package sequel

import "database/sql"

type StatementBuilder struct {
	PlaceholderFormat PlaceholderTypes
}

func New() *StatementBuilder {
	return &StatementBuilder{
		PlaceholderFormat: QuestionPlaceholderFormat,
	}
}

func (sb *StatementBuilder) SetPlaceholderFormat(placeholderType PlaceholderTypes) *StatementBuilder {
	sb.PlaceholderFormat = placeholderType
	return sb
}

// Sqlizer is an interface for objects that can be converted into SQL
type Sqlizer interface {
	ToSql() (string, []interface{}, error)
}

// BaseRunner is an interface that can execute SQL queries
type BaseRunner interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// QueryRower is an extension of BaseRunner that can also execute QueryRow
type QueryRower interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}
