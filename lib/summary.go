package lib

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func rowBuilder(message struct {
	Msg      string                     "json:\"msg\""
	Metadata struct{ details struct{} } "json:\"metadata,omitempty\""
}, level string) string {
	policyIDRegex := regexp.MustCompile("DOMI-[A-Z]*-[0-9]{3}")
	policyID := string(policyIDRegex.Find([]byte(message.Msg)))
	failureMessage := strings.Join(strings.Split(message.Msg, ": ")[1:], ": ")
	return fmt.Sprintf("| %s | %s | %s |\n", policyID, level, failureMessage)
}

// SummaryBuilder - Builds the summary for the check run.
func SummaryBuilder(conftestResults ConftestResults) (string, string) {
	summaryIntro := "**Status**: Complete\n"
	// summaryError := ""
	// if scanErr != nil {
	// 	summaryError = fmt.Sprintf("**Errors**: Are you using a custom policy repository? Check out the domi [Wiki](https://github.com/devops-kung-fu/domi/wiki) for recommendations on creating custom policy repos.\n<pre>%s</pre>", scanErr)
	// }
	summaryResultsTitle := "---\n# Results"
	summaryResultsByFile := ""
	conclusion := "success"
	log.Println(conftestResults)
	if len(conftestResults) > 0 {
		for _, result := range conftestResults {
			if len(result.Failures) > 0 || len(result.Warnings) > 0 {
				summaryResultsByFile += fmt.Sprintf("\n**%s**\n", strings.Join(strings.Split(result.Filename, "/")[4:], "/"))
				summaryResultsByFile += "| Policy | Level | Description |\n| ------ | ----- | ----------- |\n"
				if len(result.Failures) > 0 {
					for _, failure := range result.Failures {
						summaryResultsByFile += rowBuilder(failure, "Deny")
						conclusion = "failure"
					}
				}
				if len(result.Warnings) > 0 {
					for _, warning := range result.Warnings {
						summaryResultsByFile += rowBuilder(warning, "Warn")

					}
				}
			}
		}
	}
	if summaryResultsByFile == "" {
		summaryResultsByFile = "No results found. :rocket:"
	}
	if os.Getenv("AUDIT_MODE") == "1" {
		conclusion = "neutral"
	}
	summary := fmt.Sprintf("%s\n%s\n%s\n", summaryIntro, summaryResultsTitle, summaryResultsByFile)
	return summary, conclusion
}
