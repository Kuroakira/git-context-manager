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
	case "help", "-h", "--help":
		printUsage()
		return
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
	fmt.Println("gctx - Git Context Manager")
	fmt.Println()
	fmt.Println("A CLI tool that easily switches Git identity on a per-repository basis.")
	fmt.Println()
	fmt.Println("Usage: gctx <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  add <name> --email=<email> --name=<user> [--ssh-key=<path>]")
	fmt.Println("        Add a new context")
	fmt.Println()
	fmt.Println("  list")
	fmt.Println("        List available contexts")
	fmt.Println()
	fmt.Println("  use <name>")
	fmt.Println("        Apply a context to the current repository")
	fmt.Println()
	fmt.Println("  clone <name> <repository-url> [directory]")
	fmt.Println("        Clone a repository with a specific context")
	fmt.Println()
	fmt.Println("  status")
	fmt.Println("        Show current context status")
	fmt.Println()
	fmt.Println("  remove <name>")
	fmt.Println("        Remove a context")
	fmt.Println()
	fmt.Println("  help, -h, --help")
	fmt.Println("        Show this help message")
}
