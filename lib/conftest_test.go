package lib

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	fs := OSFS{}
	result := Scan(fs, "/__testdata__", "test-policy", []string{"/__testdata__/test.tf"})
	fmt.Println(result)
}
