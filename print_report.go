package main

import (
	"fmt"
	"sort"
)

type report struct {
	page   string
	visits int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
REPORT for %s
=============================
`, baseURL)
	slice := make([]report, 0, len(pages))
	for page, num := range pages {
		slice = append(slice, report{visits: num, page: page})
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].visits > slice[j].visits
	})

	for _, r := range slice {
		fmt.Printf("Found %d internal links to %s\n", r.visits, r.page)
	}
}
