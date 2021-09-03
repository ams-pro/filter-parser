package serializer_test

import (
	"testing"

	filterparser "github.com/ams-pro/filter-parser"
	"github.com/ams-pro/filter-parser/serializer"
)

func TestSerializer(t *testing.T) {

	filter := "in(accessibleUnits.unit, [5eda92fe7eed83001c78ae1c])"
	node, _ := filterparser.ParseFilter(filter)

	s := &serializer.Mongo{}

	f, err := s.Parse(node)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", f)

}
