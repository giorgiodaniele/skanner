package main

import (
	"net"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	Target string
	Port   int
}

type Result struct {
	Target string
	Port   int
	Status string
}

func scanPort(job Job, timeout time.Duration) Result {
	address := net.JoinHostPort(
		job.Target,
		strconv.Itoa(job.Port),
	)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return Result{
			Target: job.Target,
			Port:   job.Port,
			Status: "CLOSED/FILTERED",
		}
	}

	_ = conn.Close()

	return Result{
		Target: job.Target,
		Port:   job.Port,
		Status: "OPEN",
	}
}

func runScan(
	targets []string,
	ports []int,
	timeout time.Duration,
	workers int,
) []Result {
	jobs := make(chan Job)
	results := make(chan Result)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for job := range jobs {
				results <- scanPort(job, timeout)
			}
		}()
	}

	go func() {
		for _, target := range targets {
			for _, port := range ports {
				jobs <- Job{Target: target, Port: port}
			}
		}
		close(jobs)

		wg.Wait()
		close(results)
	}()

	var all []Result
	for result := range results {
		all = append(all, result)
	}

	return all
}