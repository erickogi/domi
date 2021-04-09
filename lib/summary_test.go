package lib

import (
	"fmt"
	"testing"
)

func TestRowBuilder(t *testing.T) {
	result := rowBuilder(struct {
		Msg      string                     "json:\"msg\""
		Metadata struct{ details struct{} } "json:\"metadata,omitempty\""
	}{Msg: "DOMI-TET-666: Blah Blah Blah", Metadata: struct{ details struct{} }{}}, "Warn")
	fmt.Println(result)
}
