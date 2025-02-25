package sequel

type Builder interface {
	ToSql() (string, []interface{}, error)
}
