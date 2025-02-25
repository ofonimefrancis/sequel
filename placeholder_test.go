package sequel_test

import (
	"testing"

	"github.com/ofonimefrancis/sequel"
	"github.com/stretchr/testify/assert"
)

func TestPlaceholder_QuestionFormat(t *testing.T) {
	questionFormat := sequel.Question

	s, err := questionFormat.ReplacePlaceholders("SELECT * FROM users WHERE id = ?")
	assert.NoError(t, err)

	assert.Equal(t, s, "SELECT * FROM users WHERE id = ?")
}
