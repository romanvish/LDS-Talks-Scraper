package main

import (
	"fmt"
	"os"
	"os/exec"
	// "strconv"
	// "sync"
)

func createRange(n int) []int {
	// function to use to create a slice (like python list(range(n)))
	// refer to claude.ai/chat/bb679497-dd9f-4097-a9b8-9e9dceb191ee for more
	slice := make([]int, n)
	for i := range slice {
		slice[i] = i
	}
	return slice
}

// this is currently serial code lol. but i'm going to parallelize it once i get it working.
func main() {

	how_long := os.Args[1] // leave as a string bc we're passing it to the shell script
	// TODO -- parallelize this, now that it works
	// talk_nums := createRange(how_long)

	scrape_cmd := exec.Command("sh", "scrape2.sh", how_long)
	scrape_output, scrape_err := scrape_cmd.CombinedOutput()

	if scrape_err != nil {
		fmt.Printf("Error running script: %v\n", scrape_err)
		fmt.Printf("Error output: %s\n", scrape_output)
		return
	}

	fmt.Println("starting talk processor...")

	cmd := exec.Command("python3", "talkParser2.py")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error running script: %v\n", err)
		fmt.Printf("Error output: %s\n", output)
		return
	}

	fmt.Printf("Python output:\n%s", output)
	fmt.Println("Script completed successfully!")
}
