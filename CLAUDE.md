# Project: git-context-manager (Command: gctx)

## Overview
A CLI tool that easily switches Git identity ("who to act as" - Name, Email, SSH Key) on a per-repository basis.
Aims to provide an experience similar to Python's `venv`, where settings are applied (`activate`) per repository.

## Core Goals
1.  **Identity Switching:** Modify `user.name` and `user.email` via `--local` configuration.
2.  **SSH Abstraction:** Accept SSH key path and automatically generate `core.sshCommand` configuration (no URL rewriting needed).
3.  **Persistence:** Store settings (Context) in a JSON file under the user's home directory.

## Tech Stack
* Language: Go (Golang)
* Standard Library first (prioritize standard library; use minimal external libraries only when necessary)

## Architecture & Design
* **Context Structure:**
    * Name (string) - Context identifier
    * UserName (string) - Git user.name
    * Email (string) - Git user.email
    * SSHKeyPath (string) - Path to SSH private key
* **Storage:** `~/.gctxconfig`
* **Git Interaction:** Use `os/exec` to execute `git config --local` commands.

## Commands (Specification)
* `gctx add <context-name> --email=<...> --name=<...> [--ssh-key=<...>]`: Save configuration (email and name are required).
* `gctx use <context-name>`: Apply configuration to current directory's `.git/config`.
* `gctx clone <context-name> <repository-url> [directory]`: Clone a repository using the specified context's SSH key, then apply the context.
* `gctx list`: Display list of registered configurations.
* `gctx status`: Display configuration currently applied to the repository.
* `gctx remove <context-name>`: Remove a configuration.

## Development Guidelines
* **Error Handling:** Don't suppress errors; display user-friendly error messages.
* **Code Style:** Follow `gofmt` conventions.
* **Build:** `go build -o gctx main.go`
* **Test:** `go test ./...`

## Memo for AI
* User prioritizes "engineering efficiency."
* User is still learning Go implementation; provide clear code explanations.
* Prioritize working MVP (Minimum Viable Product) over complex abstractions.
