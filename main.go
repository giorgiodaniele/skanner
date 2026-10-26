package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	targetFlag := flag.String(
		"targets", "127.0.0.1",
		"Comma-separated targets",
	)
	portFlag := flag.String(
		"ports", "22,80,443,1-1024",
		"Comma-separated ports and ranges",
	)
	timeout := flag.Duration(
		"timeout", 500*time.Millisecond,
		"Connection timeout",
	)
	workers := flag.Int(
		"workers", 100,
		"Maximum concurrent scans",
	)
	flag.Parse()

	if *workers < 1 || *timeout <= 0 {
		fmt.Fprintln(os.Stderr, "Invalid workers or timeout")
		os.Exit(1)
	}

	targets := splitTargets(*targetFlag)
	if len(targets) == 0 {
		fmt.Fprintln(os.Stderr, "No targets specified")
		os.Exit(1)
	}

	ports, err := parsePorts(*portFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	results := runScan(targets, ports, *timeout, *workers)
	printResults(results)

}
func splitTargets(input string) []string {
	seen := make(map[string]bool)
	var targets []string

	for _, target := range strings.Split(input, ",") {
		target = strings.TrimSpace(target)
		if target != "" && !seen[target] {
			targets = append(targets, target)
			seen[target] = true
		}
	}
	return targets
}