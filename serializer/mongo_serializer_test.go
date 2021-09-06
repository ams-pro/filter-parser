package serializer_test

import (
	"testing"

	filterparser "github.com/ams-pro/filter-parser"
	"github.com/ams-pro/filter-parser/serializer"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMongoSerializer(t *testing.T) {
	oid, _ := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
	testCases := []struct {
		desc string
		in   string
		out  bson.M
	}{
		{
			desc: "Parse simple eq expression",
			in:   `eq(_id, abc)`,
			out:  bson.M{"_id": "abc"},
		},
		{
			desc: "Parse simple eq expression with ObjectID",
			in:   `eq(_id, 507f1f77bcf86cd799439011)`,
			out:  bson.M{"_id": oid},
		},
		{
			desc: "Parse and expression consisting of simple eq expression",
			in:   `and(eq(_id, abc),eq(email, bcd))`,
			out:  bson.M{"_id": "abc", "email": "bcd"},
		},
		{
			desc: "Parse and expression consisting of simple eq and in expression",
			in:   `and(eq(_id, abc),in(units, [bcd,efg]))`,
			out:  bson.M{"_id": "abc", "units": bson.M{"$in": []string{"bcd", "efg"}}},
		},
		{
			desc: "Parse and expression consisting of simple eq and in expression (int)",
			in:   `and(eq(_id, abc),in(units, [123,345]))`,
			out:  bson.M{"_id": "abc", "units": bson.M{"$in": []int64{123, 345}}},
		},
		{
			desc: "Parse and expression consisting of simple eq and in expression (OID)",
			in:   `and(eq(_id, abc),in(units, [507f1f77bcf86cd799439011,507f1f77bcf86cd799439011]))`,
			out:  bson.M{"_id": "abc", "units": bson.M{"$in": []primitive.ObjectID{oid, oid}}},
		},
		{
			desc: "Parse in expression with empty slice",
			in:   "in(email, [])",
			out:  bson.M{"email": bson.M{"$in": []interface{}{}}},
		},
	}
	for _, tC := range testCases {
		s := &serializer.Mongo{}
		t.Run(tC.desc, func(t *testing.T) {
			n, _ := filterparser.ParseFilter(tC.in)
			f, _ := s.Parse(n)
			assert.Equal(t, tC.out, f)
		})
	}
}
