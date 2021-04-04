package lib

import (
	"testing"
	"regexp"
)

func TestGetUUID(t *testing.T) {
	result := getUUID()
	uuidRegex := regexp.MustCompile("[a-f0-9]{8}-([a-f0-9]{4}-){3}[a-f0-9]{12}")
	uuidMatch := uuidRegex.FindString(result)
	if uuidMatch == "" {
		t.Error()
	}
}