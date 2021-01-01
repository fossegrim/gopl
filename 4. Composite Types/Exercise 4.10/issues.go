package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	// The GitHub REST API supports requesting issues wihin certain timespans.
	// We could alternatively use this feature instead of sorting them locally.
	// Both approaches are equally valid in my opinon.
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var lessThanAMonthOldIssues []*github.Issue
	var lessThanAYearOldIssues []*github.Issue // The intersection of this and lessThanAMonthOldIssues shall be empty.
	var moreThanAYearOldIssues []*github.Issue // The intersection of this and lessThanAYearOldIssues shall be empty.

	aMonthAgo := time.Now().AddDate(0, -1, 0)
	aYearAgo := time.Now().AddDate(-1, 0, 0)

	for _, item := range result.Items {
		if item.CreatedAt.After(aMonthAgo) {
			lessThanAMonthOldIssues = append(lessThanAMonthOldIssues, item)
		} else if item.CreatedAt.After(aYearAgo) {
			lessThanAYearOldIssues = append(lessThanAYearOldIssues, item)
		} else {
			moreThanAYearOldIssues = append(moreThanAYearOldIssues, item)
		}
	}

	fmt.Printf("%d less than a month old issues:\n", len(lessThanAMonthOldIssues))
	printIssues(lessThanAMonthOldIssues)
	fmt.Printf("%d less than a year old issues:\n", len(lessThanAYearOldIssues))
	printIssues(lessThanAYearOldIssues)
	fmt.Printf("%d more than a year old issues:\n", len(moreThanAYearOldIssues))
	printIssues(moreThanAYearOldIssues)
}

func printIssues(issues []*github.Issue) {
	for _, issue := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}
}
