package lib

import (
	"fmt"
	"log"
	"os"
)

// SummaryBuilder - Builds the summary for the check run.
func SummaryBuilder(conftestResults ConftestResults) (string, string) {
	summaryIntro := "**Status**: Complete"
	// summaryError := ""
	// if scanErr != nil {
	// 	summaryError = fmt.Sprintf("**Errors**: Are you using a custom policy repository? Check out the domi [Wiki](https://github.com/devops-kung-fu/domi/wiki) for recommendations on creating custom policy repos.\n<pre>%s</pre>", scanErr)
	// }
	summaryResultsTitle := "---\n# Results"
	summaryResultsByFile := ""
	conclusion := ""
	log.Println(conftestResults)
	if len(conftestResults) > 0 {
		for _, result := range conftestResults {
			if len(result.Failures) == 0 || len(result.Warnings) == 0 {
				break
			}
			summaryResultsByFile += fmt.Sprintf("*%s*\n", result.Filename)
		}
		conclusion = "failure"
	}
	if summaryResultsByFile == "" {
		summaryResultsByFile = "No results found. :rocket:"
		conclusion = "success"
	}
	if os.Getenv("AUDIT_MODE") == "1" {
		conclusion = "neutral"
	}
	summary := fmt.Sprintf("%s\n%s\n%s\n", summaryIntro, summaryResultsTitle, summaryResultsByFile)
	return summary, conclusion
}