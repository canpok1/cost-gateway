package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryBuilder(t *testing.T) {
	t.Run("Build", func(t *testing.T) {
		t.Run("条件なし", func(t *testing.T) {
			gotSQL, gotValue := NewQueryBuilder("SELECT * FROM x").Build()

			wantSQL := "SELECT * FROM x"
			wantValue := []interface{}{}

			assert.Equal(t, wantSQL, gotSQL)

			assert.Equal(t, wantValue, gotValue)
		})

		t.Run("条件あり", func(t *testing.T) {
			gotSQL, gotValue := NewQueryBuilder("SELECT * FROM x").
				AddCondition("value1 = ?", 1).
				AddCondition("value2 = ?", "2").
				Build()

			wantSQL := "SELECT * FROM x WHERE value1 = ? AND value2 = ?"
			wantValue := []interface{}{1, "2"}

			assert.Equal(t, wantSQL, gotSQL)
			assert.Equal(t, wantValue, gotValue)
		})
	})
}
