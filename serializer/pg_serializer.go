package serializer

import (
	"strings"
)

type Postgres struct {
}

type WherableBuilder interface {
	GreaterEqualThan(field string, value interface{}) string
	GreaterThan(field string, value interface{}) string
	LessEqualThan(field string, value interface{}) string
	LessThan(field string, value interface{}) string
	NotEqual(field string, value interface{}) string
	Equal(field string, value interface{}) string
	In(field string, value ...interface{}) string
	Or(orExpr ...string) string
	And(andExpr ...string) string
}

func parsePostgres(expr []string, sb WherableBuilder) string {
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

			p1 := parsePostgres(st1, sb)
			p2 := parsePostgres(st2, sb)
			return sb.Or(p1, p2)
		}
	}

	return sb.And(where...)
}

func (*Postgres) Parse(expr []string, sb WherableBuilder) string {
	return parsePostgres(expr, sb)
}
