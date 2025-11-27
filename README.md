# git-context-manager (gctx)

A CLI tool that easily switches Git identity ("who to act as" - Name, Email, SSH Key) on a per-repository basis.
It aims to provide an experience similar to Python's `venv`, where settings are applied per repository.

## Features

- **Identity Switching**: Easily switch `user.name`, `user.email`, and `core.sshCommand` for the current repository.
- **SSH Key Management**: Automatically configures `core.sshCommand` to use a specific SSH key for the repository.
- **Persistence**: Saves your contexts in `~/.gctxconfig`.

## Installation

Currently, you can build the tool from source.

### Prerequisites

- Go (Golang) installed

### Build

```bash
git clone https://github.com/Kuroakira/git-context-manager.git
cd git-context-manager
go build -o gctx main.go
```

You can then move the `gctx` binary to a directory in your `$PATH` (e.g., `/usr/local/bin`) or use it directly.

## Usage

### 1. Register a Context

Save a new identity configuration.

```bash
gctx add <name> --email=<email> --name=<user_name> [--ssh-key=<path_to_private_key>]
```

Example:
```bash
gctx add personal --email=me@example.com --name="My Name" --ssh-key=~/.ssh/id_rsa_personal
```

### 2. List Contexts

Show all registered contexts.

```bash
gctx list
```

### 3. Apply Context

Apply a context to the current git repository. This modifies the local `.git/config`.

```bash
cd /path/to/your/repo
gctx use <name>
```

Example:
```bash
gctx use personal
```

### 4. Check Status

Show the currently applied configuration for the repository.

```bash
gctx status
```

### 5. Remove Context

Remove a registered context.

```bash
gctx remove <name>
```

## Configuration

Configuration is stored in `~/.gctxconfig` in JSON format.
