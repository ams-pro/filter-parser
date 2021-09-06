package serializer_test

import (
	"testing"

	filterparser "github.com/ams-pro/filter-parser"
	"github.com/ams-pro/filter-parser/serializer"
	"github.com/huandu/go-sqlbuilder"
	"github.com/stretchr/testify/assert"
)

func TestPostgresoSerializer(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		out  string
	}{
		{
			desc: "Parse simple eq expression",
			in:   `eq(_id, abc)`,
			out:  "(_id = $0)",
		},
		{
			desc: "Parse simple eq expression with ObjectID",
			in:   `eq(_id, 507f1f77bcf86cd799439011)`,
			out:  "(_id = $0)",
		},
		{
			desc: "Parse and expression consisting of simple eq expression",
			in:   `and(eq(_id, abc),eq(email, bcd))`,
			out:  "(_id = $0 AND email = $1)",
		},
		{
			desc: "Parse and expression consisting of simple eq and in expression",
			in:   `and(eq(_id, abc),in(units, [bcd,efg]))`,
			out:  "(_id = $0 AND units IN ($1))",
		},
		{
			desc: "Parse in expression with empty slice",
			in:   "in(email, [])",
			out:  "(email IN ($0))",
		},
		{
			desc: "Parse or expression consisting of simple eq expression",
			in:   `or(eq(_id, abc),eq(email, bcd))`,
			out:  "((_id = $0) OR (email = $1))",
		},
	}
	for _, tC := range testCases {
		s := &serializer.Postgres{}
		t.Run(tC.desc, func(t *testing.T) {
			n, _ := filterparser.ParseFilter(tC.in)
			sb := sqlbuilder.NewSelectBuilder()
			f := s.Parse(n.Inorder(), sb)
			assert.Equal(t, tC.out, f)
		})
	}
}
