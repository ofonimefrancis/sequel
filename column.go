package sequel

// Column represents a SQL column
type Column struct {
	name string
}

func newColumn(name string) Column {
	return Column{name: name}
}

func (c Column) ToSql() (string, []interface{}, error) {
	return c.name, nil, nil
}
