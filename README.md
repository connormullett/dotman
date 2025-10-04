# dotman

[![Go Version](https://img.shields.io/badge/Go-1.21.5-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

A simple yet powerful CLI tool to manage your dotfiles using Git as the backend. Keep your configuration files synchronized across multiple machines with ease.

## 📋 Table of Contents

- [Features](#-features)
- [Installation](#-installation)
- [Quick Start](#-quick-start)
- [Usage](#-usage)
  - [Initialize](#initialize)
  - [Add Dotfiles](#add-dotfiles)
  - [List Managed Files](#list-managed-files)
  - [Remove Dotfiles](#remove-dotfiles)
  - [Push Changes](#push-changes)
  - [Sync](#sync)
  - [Health Check](#health-check)
- [How It Works](#-how-it-works)
- [Configuration](#-configuration)
- [Development](#-development)
- [Contributing](#-contributing)
- [License](#-license)

## ✨ Features

- **Git-backed**: Uses Git as the storage backend for version control and history
- **Symlink Management**: Automatically creates and manages symlinks to keep files in their expected locations
- **Cross-machine Sync**: Easily synchronize your dotfiles across multiple machines
- **Health Checks**: Detect broken symlinks and configuration issues
- **Simple CLI**: Intuitive command-line interface built with Cobra
- **Safe Operations**: Validates operations to prevent accidental data loss

## 📦 Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/connormullett/dotman.git
cd dotman

# Build and install
make build

# Or install directly to GOPATH/bin
make install
```

### Using Go Install

```bash
go install github.com/dotman@latest
```

## 🚀 Quick Start

1. **Initialize dotman with your remote repository:**

```bash
dotman init git@github.com:yourusername/dotfiles.git
```

2. **Add your first dotfile:**

```bash
dotman add ~/.bashrc
```

3. **Push to your remote repository:**

```bash
dotman push
```

4. **On a new machine, initialize and sync:**

```bash
dotman init git@github.com:yourusername/dotfiles.git
dotman sync
```

## 📖 Usage

### Initialize

Initialize dotman with a remote Git repository:

```bash
dotman init <remote-url>
```

**Example:**
```bash
dotman init git@github.com:yourusername/dotfiles.git
```

This will:
- Create a `.dotman` directory in your home folder
- Initialize a Git repository
- Add the specified remote as `origin`

### Add Dotfiles

Add a dotfile to dotman management:

```bash
dotman add <file-path>
```

**Example:**
```bash
dotman add ~/.vimrc
dotman add ~/.zshrc
dotman add ~/.gitconfig
```

When you add a file, dotman will:
1. Move the file to `~/.dotman/`
2. Create a symlink at the original location pointing to the new location
3. Stage and commit the file to Git

**Note:** If a file is already a symlink, it will be skipped automatically.

### List Managed Files

View all dotfiles currently managed by dotman:

```bash
dotman list
```

This displays all files in your `~/.dotman` repository.

### Remove Dotfiles

Remove a dotfile from dotman management:

```bash
dotman remove <filename>
```

**Example:**
```bash
dotman remove .vimrc
```

This will:
1. Remove the symlink from your home directory
2. Move the actual file back to its original location
3. Commit the removal to Git

**Safety:** The command validates that the symlink points to the dotman repository before removal to prevent data loss.

### Push Changes

Push your dotfiles to the remote repository:

```bash
dotman push
```

**Options:**
- `-f, --force`: Force push changes (use with caution)

```bash
dotman push --force
```

If there are uncommitted changes, dotman will automatically create a commit with the message "auto-commit before push" before pushing.

### Sync

Synchronize your dotfiles from the remote repository:

```bash
dotman sync
```

This command:
- Checks if the repository has uncommitted changes
- Pulls the latest changes from the remote repository
- **Important:** Will fail if there are uncommitted changes to prevent merge conflicts

**Pro Tip:** Add this to your shell's startup file (e.g., `.zshrc` or `.bashrc`) to automatically sync on shell startup:

```bash
# Auto-sync dotfiles on shell startup
dotman sync --quiet
```

### Health Check

Check the health of your dotman setup:

```bash
dotman doctor
```

This command:
- Verifies all managed dotfiles
- Detects broken symlinks
- Reports any configuration issues
- Checks for circular symlinks

## 🔧 How It Works

dotman uses a clever symlink-based approach:

1. **Storage**: All your dotfiles are stored in `~/.dotman/`, which is a Git repository
2. **Symlinks**: Your actual dotfiles (e.g., `~/.vimrc`) are symlinks pointing to files in `~/.dotman/`
3. **Version Control**: Changes to your dotfiles are tracked in the Git repository
4. **Synchronization**: You can push/pull changes to/from a remote repository

### Directory Structure

```
~/.dotman/
├── .git/              # Git repository
├── .vimrc             # Your actual .vimrc file
├── .zshrc             # Your actual .zshrc file
├── .gitconfig         # Your actual .gitconfig file
└── ...                # Other dotfiles

~/
├── .vimrc -> ~/.dotman/.vimrc        # Symlink
├── .zshrc -> ~/.dotman/.zshrc        # Symlink
├── .gitconfig -> ~/.dotman/.gitconfig # Symlink
└── ...
```

## ⚙️ Configuration

dotman stores its configuration in a platform-specific location:

- **macOS**: `~/Library/Application Support/dotman/config.json`
- **Linux**: `~/.config/dotman/config.json`
- **Windows**: `%APPDATA%/dotman/config.json`

The configuration file contains:

```json
{
  "path": "/Users/username/.dotman"
}
```

This path can be customized to change where dotman stores your dotfiles.

## 🛠️ Development

### Prerequisites

- Go 1.21.5 or later
- Git

### Building

```bash
# Install dependencies
make deps

# Build the project
make build

# Run tests
make test

# Clean build artifacts
make clean
```

### Project Structure

```
dotman/
├── main.go              # Entry point
├── cmd/
│   └── root.go          # CLI command definitions
├── subcommands/
│   ├── add.go           # Add command implementation
│   ├── config.go        # Configuration management
│   ├── doctor.go        # Health check implementation
│   ├── git.go           # Git operations
│   ├── init.go          # Initialize command
│   ├── list.go          # List command
│   ├── push.go          # Push command
│   ├── remove.go        # Remove command
│   └── sync.go          # Sync command
├── go.mod               # Go module definition
├── go.sum               # Dependency checksums
└── Makefile             # Build automation
```

### Dependencies

- **[cobra](https://github.com/spf13/cobra)**: CLI framework
- **[configdir](https://github.com/kirsle/configdir)**: Cross-platform config directory management

## 🤝 Contributing

Contributions are welcome! Here are some ways you can contribute:

1. Report bugs and feature requests
2. Submit pull requests
3. Improve documentation
4. Share your dotfiles setup

### Development Workflow

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests (`make test`)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to your branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) CLI framework
- Inspired by the need for simple dotfile management

## 📧 Contact

Connor Mullett - [@connormullett](https://github.com/connormullett)

Project Link: [https://github.com/connormullett/dotman](https://github.com/connormullett/dotman)

---

**Note**: Always backup your dotfiles before using any dotfile management tool for the first time!
