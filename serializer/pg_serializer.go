package serializer

import (
	"strings"

	"github.com/huandu/go-sqlbuilder"
)

func Parse(expr []string, sb *sqlbuilder.SelectBuilder) string {
	where := []string{}
	for i, e := range expr {
		switch e {
		case "ge":
			lhs := strings.TrimSpace(expr[i-1])
			rhs := strings.TrimSpace(expr[i+1])
			gt := sb.GreaterEqualThan(lhs, rhs)
			where = append(where, gt)
		case "gt":
			lhs := strings.TrimSpace(expr[i-1])
			rhs := strings.TrimSpace(expr[i+1])
			gt := sb.GreaterThan(lhs, rhs)
			where = append(where, gt)
		case "le":
			lhs := strings.TrimSpace(expr[i-1])
			rhs := strings.TrimSpace(expr[i+1])
			lt := sb.LessEqualThan(lhs, rhs)
			where = append(where, lt)
		case "lt":
			lhs := strings.TrimSpace(expr[i-1])
			rhs := strings.TrimSpace(expr[i+1])
			lt := sb.LessThan(lhs, rhs)
			where = append(where, lt)
		case "ne":
			lhs := strings.TrimSpace(expr[i-1])
			rhs := strings.TrimSpace(expr[i+1])
			eq := sb.NotEqual(lhs, rhs)
			where = append(where, eq)
		case "eq":
			lhs := strings.TrimSpace(expr[i-1])
			rhs := strings.TrimSpace(expr[i+1])
			eq := sb.Equal(lhs, rhs)
			where = append(where, eq)
		case "in":
			lhs := strings.TrimSpace(expr[i-1])
			rhs := strings.TrimSpace(expr[i+1])
			e := strings.Trim(rhs, "[]")
			in := sb.In(lhs, e)
			where = append(where, in)
		case "or":
			st1 := expr[:i]
			st2 := expr[i+1:]

			p1 := Parse(st1, sb)
			p2 := Parse(st2, sb)
			return sb.Or(p1, p2)
		}
	}

	return sb.And(where...)
}
