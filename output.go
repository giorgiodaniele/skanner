package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

func printResults(results []Result) {
	sort.Slice(results, func(i, j int) bool {
		if results[i].Target != results[j].Target {
			return results[i].Target < results[j].Target
		}
		return results[i].Port < results[j].Port
	})

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, "TARGET\tPORT\tSERVICE\tSTATUS")
	fmt.Fprintln(w, "------\t----\t-------\t------")

	for _, r := range results {
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\n",
			r.Target,
			r.Port,
			serviceName(r.Port),
			r.Status,
		)
	}

	_ = w.Flush()
}