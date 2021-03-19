package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
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
	Warnings	[]struct {
		Msg			string	`json:"msg"`
		Metadata	struct {
			details	struct {

			}
		} `json:"metadata,omitempty"`
	} `json:"warnings,omitempty"`
}

// Scan - Use conftest to Scan discovered files.
func Scan(fs fileSystem, policyID string, files []string) (ConftestResults, error) {
	policyIDPath := fmt.Sprintf("/tmp/%s", policyID)
	policyPaths, policyPathsError := FindFiles(fs, policyIDPath, "policy$")
	if policyPathsError != nil {
		return nil, policyPathsError
	}
	var policyPath string = ""
	for _, policyPathCandidate := range policyPaths {
		if len(strings.Split(policyPathCandidate, "/")) == 5 {
			policyPath = policyPathCandidate
			break
		}
	}
	cmd := exec.Command("conftest", "test", "--all-namespaces", "-o", "json", "-p", policyPath)
	cmd.Args = append(cmd.Args, files...)
	output, outputErr := cmd.Output()
	if outputErr != nil {
		return nil, outputErr
	}
	log.Println(output)
	conftestResults := ConftestResults{}
	json.Unmarshal(output, &conftestResults)
	return conftestResults, nil
}