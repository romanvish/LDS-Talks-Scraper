package main

import (
	"fmt"
	"os/exec"
	// "os"
)

func main() {
	fmt.Println("starting talk processor...")

	cmd := exec.Command("python3", "talkParser.py")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error running script: %v\n", err)
		fmt.Printf("Error output: %s\n", output)
		return
	}

	fmt.Printf("Python output:\n%s", output)
	fmt.Println("Script completed successfully!")
}
