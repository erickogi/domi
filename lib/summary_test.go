package lib

import (
	"fmt"
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

func TestSummaryBuilder(t *testing.T) {
	conftestResults := &ConftestResults{
		{
			Filename:  "fake.file",
			Successes: 124,
			Failures: []struct {
				Msg      string `json:"msg"`
				Metadata struct {
					details struct{}
				} `json:"metadata,omitempty"`
			}{
				{
					Msg: "DOMI-TEST-001: Test Policy",
				},
				{
					Msg: "DOMI-TEST-002: Test Policy",
				},
			},
			Warnings: []struct {
				Msg      string `json:"msg"`
				Metadata struct {
					details struct{}
				} `json:"metadata,omitempty"`
			}{
				{
					Msg: "DOMI-TEST-003: Test Policy",
				},
				{
					Msg: "DOMI-TEST-004: Test Policy",
				},
			},
		},
	}
	summary, conclusion := SummaryBuilder(conftestResults)
	fmt.Println(summary)
	fmt.Println(conclusion)
}
