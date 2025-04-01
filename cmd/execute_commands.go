package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Status   int
	Duration time.Duration
}

func RunTest(url string, totalRequests, concurrency int) {
	start := time.Now()

	var wg sync.WaitGroup
	var progressWG sync.WaitGroup

	sem := make(chan struct{}, concurrency)
	results := make(chan Result, totalRequests)
	progress := make(chan int, totalRequests)

	progressWG.Add(1)
	go func() {
		defer progressWG.Done()
		barLength := 30
		for i := 1; i <= totalRequests; i++ {
			<-progress
			percent := (i * 100) / totalRequests
			filledLength := (i * barLength) / totalRequests
			bar := ""
			for j := 0; j < barLength; j++ {
				if j < filledLength {
					bar += "█"
				} else {
					bar += "-"
				}
			}
			fmt.Printf("\r[%s] %3d%%", bar, percent)
		}
		fmt.Print("\r[██████████████████████████████] 100%\n\n")
	}()

	for range totalRequests {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			startTime := time.Now()
			resp, err := http.Get(url)
			elapsed := time.Since(startTime)

			if err != nil {
				results <- Result{Status: 0, Duration: elapsed}
			} else {
				results <- Result{Status: resp.StatusCode, Duration: elapsed}
				resp.Body.Close()
			}
			<-sem
			progress <- 1
		}()
	}

	wg.Wait()
	close(results)
	progressWG.Wait()

	duration := time.Since(start)
	generateReport(results, totalRequests, duration)
}


func generateReport(results chan Result, total int, duration time.Duration) {
	statusCount := make(map[int]int)
	var totalRequestTime time.Duration

	for res := range results {
		statusCount[res.Status]++
		totalRequestTime += res.Duration
	}

	avg := totalRequestTime / time.Duration(total)

	fmt.Println("=== RELATÓRIO DE TESTE ===")
	fmt.Printf("Tempo total: %v\n", duration)
	fmt.Printf("Tempo médio por request: %v\n", avg)
	fmt.Printf("Requests feitos: %d\n", total)
	fmt.Printf("Status 200: %d\n", statusCount[200])
	for code, count := range statusCount {
		if code != 200 {
			fmt.Printf("Status %d: %d\n", code, count)
		}
	}
}
