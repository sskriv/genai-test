package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()
	// The client gets the API key from the environment variable `GEMINI_API_KEY`.
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ask questions or 'quit' to exit.")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("prompt: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			return
		}

		input = strings.TrimSpace(input)
		input = strings.ToLower(input)

		switch input {
		case "quit":
			fmt.Println("Goodbye!")
			return
		default:
			result, err := client.Models.GenerateContent(
				ctx,
				"gemini-2.5-flash",
				genai.Text(input),
				nil,
			)
			if err != nil {
				fmt.Println("Error request: ", err)
				fmt.Println("Program terminated.")
				return
			}
			fmt.Println("gemini: ", result.Text())
		}
	}
}
