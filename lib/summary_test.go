package lib

import (
	"testing"
)

func TestRowBuilder(t *testing.T) {
	result := rowBuilder(struct {
		Msg      string                     "json:\"msg\""
		Metadata struct{ details struct{} } "json:\"metadata,omitempty\""
	}{Msg: "DOMI-TEST-666: Blah Blah Blah", Metadata: struct{ details struct{} }{}}, "Warn")
	if result == "" {
		t.Error()
	}
}
