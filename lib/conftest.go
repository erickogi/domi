package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// ConftestResults - Holds conftest results
type ConftestResults []struct {
	Filename 	string	`json:"filename"`
	Successes	int		`json:"successes,omitempty"`
	Failures	[]struct {
		Msg			string	`json:"msg"`
		Metadata	struct {
			details	struct {

			}
		} `json:"metadata,omitempty"`
	} `json:"failures,omitempty"`
}

// Scan - Use conftest to Scan discovered files.
func Scan(policyID string, files []string) (ConftestResults, error) {
	conftestExec, lookpathErr := exec.LookPath("conftest")
	if lookpathErr != nil {
		return nil, lookpathErr
	}
	policyPath := fmt.Sprintf("/tmp/%s/policies", policyID)
	arguments := []string{conftestExec, "test", "--all-namespaces", "-o", "json", "-p", policyPath}
	arguments = append(arguments, files...)
	cmd := &exec.Cmd{
		Path:	conftestExec,
		Args:	[]string{conftestExec, "test", "--all-namespaces", "-o", "json"},
	}
	var output []byte
	var outputErr error
	if output, outputErr = cmd.Output(); outputErr != nil {
		return nil, outputErr
	} 
	conftestResults := ConftestResults{}
	json.Unmarshal(output, &conftestResults)
	return conftestResults, nil
}