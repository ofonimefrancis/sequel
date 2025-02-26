package sequel

type SelectBuilder struct {
	placeholderFormat PlaceholderTypes
	runWith           BaseRunner
	prefixes          []Sqlizer
	suffixes          []Sqlizer
	options           []string
	columns           []Sqlizer
	from              Sqlizer
	joins             []Sqlizer
	where             []Sqlizer
	groupBy           []string
	orderBy           []Sqlizer
	having            []Sqlizer
	limit             string
	offsets           string
}

func (sb *StatementBuilder) Select(columns ...string) *SelectBuilder {
	builder := &SelectBuilder{
		placeholderFormat: sb.PlaceholderFormat,
	}

	for _, column := range columns {
		builder.columns = append(builder.columns, newColumn(column))
	}

	return builder
}

func (sb *SelectBuilder) From(from string) *SelectBuilder {
	sb.from = newTable(from)
	return sb
}

// Join adds a JOIN clause to the query
func (sb *SelectBuilder) Join(join string) *SelectBuilder {
	sb.joins = append(sb.joins, newJoin("JOIN", join))
	return sb
}

// LeftJoin adds a LEFT JOIN clause to the query
func (sb *SelectBuilder) LeftJoin(join string) *SelectBuilder {
	sb.joins = append(sb.joins, newJoin("LEFT JOIN", join))
	return sb
}

// RightJoin adds a RIGHT JOIN clause to the query
func (sb *SelectBuilder) RightJoin(join string) *SelectBuilder {
	sb.joins = append(sb.joins, newJoin("RIGHT JOIN", join))
	return sb
}
