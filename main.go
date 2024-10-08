package main

import (
	"fmt"
	"os"

	"github.com/running-go/aicli/cmd"
)

func main() {

	geminiAPIKey := os.Getenv("GEMINI_API_KEY")

	if geminiAPIKey == "" {
		fmt.Println("Please set the GEMINI_API_KEY environment variable. Check the README for more information.")
		return
	}

	//cohereAPIKey := os.Getenv("COHERE_API_KEY")
	//
	//if cohereAPIKey == "" {
	//	fmt.Println("Please set the COHERE_API_KEY environment variable. Check the README for more information.")
	//	return
	//}

	// Execute the root command
	cmd.Execute()

}
