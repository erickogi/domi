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
func Scan(fs fileSystem, policyID string, files []string) (ConftestResults, error) {
	conftestExec, lookpathErr := exec.LookPath("conftest")
	if lookpathErr != nil {
		return nil, lookpathErr
	}
	policyIDPath := fmt.Sprintf("/tmp/%s", policyID)
	policyPaths, policyPathsError := FindFiles(fs, policyIDPath, "policy$")
	if policyPathsError != nil {
		return nil, policyPathsError
	}
	log.Println(policyPaths)
	arguments := []string{"test", "--all-namespaces", "-o", "json", "-p", policyPaths[0]}
	arguments = append(arguments, files...)
	cmd := &exec.Cmd{
		Path:	conftestExec,
		Args:	arguments,
	}
	var output []byte
	var outputErr error
	if output, outputErr = cmd.Output(); outputErr != nil {
		return nil, outputErr
	}
	log.Println(output)
	conftestResults := ConftestResults{}
	json.Unmarshal(output, &conftestResults)
	return conftestResults, nil
}