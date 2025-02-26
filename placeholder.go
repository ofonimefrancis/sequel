package sequel

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	QuestionPlaceholderFormat = questionFormat{}
	DolarPlaceholderFormat    = dollarFormat{}
	ColonPlaceholderFormat    = colonFormat{}
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
	return findAndReplacePlaceholder(sql, d.debugPlaceholder())
}

func (d dollarFormat) debugPlaceholder() string {
	return "$"
}

type colonFormat struct{}

func (n colonFormat) ReplacePlaceholders(sql string) (string, error) {
	return findAndReplacePlaceholder(sql, n.debugPlaceholder())
}

func (n colonFormat) debugPlaceholder() string {
	return ":"
}

func findAndReplacePlaceholder(sql string, prefix string) (string, error) {
	var buf = new(bytes.Buffer)
	i := 0

	for {
		p := strings.Index(sql, "?")
		if p == -1 {
			break
		}

		i++
		buf.WriteString(sql[:p])
		fmt.Fprintf(buf, "%s%d", prefix, i)
		sql = sql[p+1:]
	}

	buf.WriteString(sql)

	return buf.String(), nil
}
