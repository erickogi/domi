package lib

import (
	// "fmt"
	"os"
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
	conftestResults := ConftestResults{
		{
			Filename:  "/tmp/deadebeef-dead-beef-dead-beefdeadbeef/repo-being-scanned/fake.file",
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
	os.Setenv("AUDIT_MODE", "0")
	summary, conclusion := SummaryBuilder(conftestResults)
	if summary == "" {
		t.Error()
	}
	if conclusion != "failure" {
		t.Error()
	}
	os.Setenv("AUDIT_MODE", "")
}

func TestSummaryBuilderSuccessNeutral(t *testing.T) {
	conftestResults := ConftestResults{
		{
			Filename:  "/tmp/deadebeef-dead-beef-dead-beefdeadbeef/repo-being-scanned/fake.file",
			Successes: 124,
		},
	}
	os.Setenv("AUDIT_MODE", "1")
	summary, conclusion := SummaryBuilder(conftestResults)
	if summary == "" {
		t.Error()
	}
	if conclusion != "neutral" {
		t.Error()
	}
	os.Setenv("AUDIT_MODE", "")
}
