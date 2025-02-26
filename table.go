package sequel

// Table is a struct that represents a table in a database.
type Table struct {
	name string
}

func (t *Table) Name() string {
	return t.name
}

func newTable(name string) Table {
	return Table{name: name}
}

func (t Table) ToSql() (string, []interface{}, error) {
	return t.name, nil, nil
}
