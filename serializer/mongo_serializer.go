package serializer

import (
	"errors"
	"strconv"
	"strings"

	filterparser "github.com/ams-pro/filter-parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mongo struct {
}

func isNotEmpty(in string) bool {
	return len(strings.Trim(in, " ")) > 0
}

func parseSliceFromString(in string) interface{} {
	slice := strings.Split(in[2:len(in)-1], ",")

	if len(slice) == 0 {
		return nil
	}

	if _, err := strconv.ParseInt(slice[0], 10, 64); err == nil {
		var returnValue []int64
		for _, item := range slice {
			j, _ := strconv.ParseInt(item, 10, 64)
			returnValue = append(returnValue, j)
		}
		return returnValue
	}

	if primitive.IsValidObjectID(slice[0]) {
		var returnValue []primitive.ObjectID
		for _, item := range slice {
			oid, _ := primitive.ObjectIDFromHex(item)
			returnValue = append(returnValue, oid)
		}
		return returnValue
	}

	return slice
}

var ErrInvalidFilterValue = errors.New("invalid filter value")

func parseMongo(filter bson.M, tree *filterparser.Node) (bson.M, error) {

	switch tree.Token {
	case "in":
		if !isNotEmpty(tree.Left.Token) || !isNotEmpty(tree.Right.Token) {
			return nil, ErrInvalidFilterValue
		}
		slice := parseSliceFromString(tree.Right.Token)
		filter["$in"] = bson.M{tree.Left.Token: slice}
	}
	return filter, nil
}

func (*Mongo) Parse(tree *filterparser.Node) (bson.M, error) {
	return parseMongo(bson.M{}, tree)
}
