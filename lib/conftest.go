package lib

import (
	"os/exec"

	"github.com/zclconf/go-cty/cty/msgpack"
	"google.golang.org/grpc/metadata"
)

// ConftestResults - Holds conftest results
type ConftestResults []struct {
	filename 	string	`json:"filename"`
	successes	int		`json:"successes,omitempty"`
	failures	[]struct {
		msg			string	`json:"msg"`
		metadata	struct {
			details	struct {

			} `json:"metadata,omitempty"`
		} `json:"failures,omitempty"`
	}
}

// Scan - Use conftest to Scan discovered files.
func Scan(policyPaths []string, files []string) (ConftestResults, error) {
	fs := OSFS{}
	policyRepoFile, policyRepoError := DownloadFile(fs, policyRepo)
	conftestExec, lookpathErr := exec.LookPath("conftest")
	if lookpathErr != nil {
		return nil, lookpathErr
	}
	var policyPathArguments []string
	for _, policyPath := range policyPaths {
		policyPathArguments = append(policyPathArguments, "-p")
		policyPathArguments = append(policyPathArguments, policyPath)
	}
	arguments := []string{conftestExec, "test", "--all-namespaces", "-o", "json"}
	arguments = append(arguments, policyPathArguments...)
	arguments = append(arguments, files...)
	for _, file := range files {
		cmd := &exec.Cmd{
			Path:	conftestExec,
			Args:	[]string{conftestExec, "test", "--all-namespaces", "-o", "json"}
		}
	}
	
}