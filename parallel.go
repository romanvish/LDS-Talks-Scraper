package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"sync"
)

func processRange(start, end int, wg *sync.WaitGroup) {
	// helped with https://claude.ai/chat/62bf84ff-4df6-4686-99ab-1aaba792b528
	defer wg.Done()

	fmt.Printf("Processing range %d to %d\n (inclusive)", start, end)

	// Convert integers to strings for the shell script
	startStr := strconv.Itoa(start)
	endStr := strconv.Itoa(end)

	// hit api for raw html files
	scrape_cmd := exec.Command("sh", "scrape2.sh", startStr, endStr)
	scrape_output, scrape_err := scrape_cmd.CombinedOutput()
	if scrape_err != nil {
		fmt.Printf("Error running scrape script for range %d-%d: %v\n", start, end, scrape_err)
		fmt.Printf("Scrape error output: %s\n", scrape_output)
		return
	}

	fmt.Printf("Scraping completed for range %d-%d, starting talk processor...\n", start, end)

	// run python processor (html to text)
	cmd := exec.Command("python3", "talkParser2.py", startStr, endStr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running Python script for range %d-%d: %v\n", start, end, err)
		fmt.Printf("Python error output: %s\n", output)
		return
	}

	fmt.Printf("Range %d-%d completed successfully!\n", start, end)
}

func main() {

	if len(os.Args) < 4 {
		fmt.Println("Usage: go run script.go <start_position> <end_position> <n_threads>")
		return
	}

	// these are the global start/end positions of the range
	// the shell script will take start/end positions that are from within the interval
	// we'll divide the interval by the number of threads and pass these numbers to the shell script as start/stop params
	// arguments to the shell script need to be strings (making an int copy for some other spots where that is useful)
	start_position := os.Args[1]
	start_position_int, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error parsing number of threads: %v\n", err)
		return
	}

	end_position := os.Args[2]
	end_position_int, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error parsing number of threads: %v\n", err)
		return
	}

	fmt.Println(start_position)
	fmt.Println(end_position)
	n_threads, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Error parsing number of threads: %v\n", err)
		return
	}

	fmt.Printf("Processing documents %d to %d with %d threads\n", start_position, end_position, n_threads)

	total_docs := end_position_int - start_position_int + 1
	docs_per_thread := total_docs / n_threads
	remainder := total_docs % n_threads

	var wg sync.WaitGroup

	current_start := start_position_int

	for i := 0; i < n_threads; i++ {
		current_end := current_start + docs_per_thread - 1

		// add remainder to the last thread
		if i == n_threads-1 {
			current_end += remainder
		}

		// check to make sure we stay within the global end position
		if current_end > end_position_int {
			current_end = end_position_int
		}

		wg.Add(1)
		go processRange(current_start, current_end, &wg)

		// Update start position for next thread
		current_start = current_end + 1
	}
	wg.Wait()
	fmt.Println("All processing completed successfully!")
}
