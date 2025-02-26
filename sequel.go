package sequel

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
