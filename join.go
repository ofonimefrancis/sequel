package sequel

import "fmt"

type Join struct {
	joinType string
	clause   string
}

func newJoin(joinType, clause string) Join {
	return Join{joinType: joinType, clause: clause}
}

func (j Join) ToSql() (string, []interface{}, error) {
	return fmt.Sprintf("%s %s", j.joinType, j.clause), nil, nil
}
