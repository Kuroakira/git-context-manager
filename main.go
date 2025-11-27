package main

import (
	"fmt"
	"os"

	"github.com/Kuroakira/git-context-manager/cmd"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	var err error
	switch command {
	case "add":
		err = cmd.Add(args)
	case "list":
		err = cmd.List(args)
	case "use":
		err = cmd.Use(args)
	case "status":
		err = cmd.Status(args)
	case "remove":
		err = cmd.Remove(args)
	case "clone":
		err = cmd.Clone(args)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: gctx <command> [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  add     Add a new context")
	fmt.Println("  list    List available contexts")
	fmt.Println("  use     Apply a context to the current repository")
	fmt.Println("  clone   Clone a repository with a specific context")
	fmt.Println("  status  Show current context status")
	fmt.Println("  remove  Remove a context")
}
