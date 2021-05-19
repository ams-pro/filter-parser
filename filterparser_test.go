package filterparser_test

import (
	"encoding/json"
	"testing"

	filterparser "github.com/ams-pro/filter-parser"
)

func TestParseFilter(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		out  *filterparser.Node
	}{
		{
			desc: "should parse basic filter",
			in:   "and(gt(id,40),lt(id,60))",
			out: &filterparser.Node{
				Token: "and",
				Left: &filterparser.Node{
					Token: "gt",
					Left: &filterparser.Node{
						Token: "id",
					},
					Right: &filterparser.Node{
						Token: "40",
					},
				},
				Right: &filterparser.Node{
					Token: "lt",
					Left: &filterparser.Node{
						Token: "id",
					},
					Right: &filterparser.Node{
						Token: "60",
					},
				},
			},
		},
		{
			desc: "should parse array like syntax",
			in:   "and(gt(id,40),in(user,[60,70,80,90]))",
			out: &filterparser.Node{
				Token: "and",
				Left: &filterparser.Node{
					Token: "gt",
					Left: &filterparser.Node{
						Token: "id",
					},
					Right: &filterparser.Node{
						Token: "40",
					},
				},
				Right: &filterparser.Node{
					Token: "in",
					Left: &filterparser.Node{
						Token: "user",
					},
					Right: &filterparser.Node{
						Token: "[60,70,80,90]",
					},
				},
			},
		},
		{
			desc: "should parse simple one level filter",
			in:   "gt(id,40)",
			out: &filterparser.Node{

				Token: "gt",
				Left: &filterparser.Node{
					Token: "id",
				},
				Right: &filterparser.Node{
					Token: "40",
				},
			},
		},
		{
			desc: "should parse \"not\" filter",
			in:   "not(gt(id,40))",
			out: &filterparser.Node{
				Token: "not",
				Left: &filterparser.Node{

					Token: "gt",
					Left: &filterparser.Node{
						Token: "id",
					},
					Right: &filterparser.Node{
						Token: "40",
					},
				},
			},
		},
		{
			desc: "should parse \"not\" filter with array",
			in:   "not(in(id,[40,50,60]))",
			out: &filterparser.Node{
				Token: "not",
				Left: &filterparser.Node{
					Token: "in",
					Left: &filterparser.Node{
						Token: "id",
					},
					Right: &filterparser.Node{
						Token: "[40,50,60]",
					},
				},
			},
		},
		{
			desc: "should parse string literals",
			in:   "like(name,\"Matthias\"))",
			out: &filterparser.Node{

				Token: "like",
				Left: &filterparser.Node{
					Token: "name",
				},
				Right: &filterparser.Node{
					Token: "\"Matthias\"",
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tree := filterparser.ParseFilter(tC.in)

			b1, err := json.Marshal(tree)

			if err != nil {
				t.Fatal(err)
			}

			b2, err := json.Marshal(tC.out)

			if err != nil {
				t.Fatal(err)
			}

			if string(b1) != string(b2) {
				t.Fatal("not equal")
			}

		})
	}
}
