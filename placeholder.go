package sequel

var (
	Question = questionFormat{}
	Dolar    = dollarFormat{}
)

type PlaceholderTypes interface {
	ReplacePlaceholders(sql string) (string, error)
}

type placeholderDebugger interface {
	debugPlaceholder() string
}

type questionFormat struct{}

func (q questionFormat) ReplacePlaceholders(sql string) (string, error) {
	return sql, nil
}

func (q questionFormat) debugPlaceholder() string {
	return "?"
}

type dollarFormat struct{}

func (d dollarFormat) ReplacePlaceholders(sql string) (string, error) {
	// TODO: find and replace dollar sign positional placeholders
	return sql, nil
}

func (d dollarFormat) debugPlaceholder() string {
	return "$"
}
