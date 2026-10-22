# Gator - Blog Aggregator CLI

Gator is a command-line RSS feed aggregator built with Go and PostgreSQL. It allows users to register, track RSS feeds, and aggregate posts across the web.

## Features

- **Multi-user support**: Register and switch between different users.
- **Feed management**: Add RSS feeds linked to the creator's account and inspect all stored feeds.
- **Persistent storage**: Backed by PostgreSQL using migrations and type-safe queries generated with SQLC.
- **Feed fetching**: Parse and decode XML RSS feeds directly from the web.

## Prerequisites

- [Go](https://go.dev/) (1.20+)
- [PostgreSQL](https://www.postgresql.org/)
- [sqlc](https://sqlc.dev/) (for generating database access code)
- [goose](https://github.com/pressly/goose) (for database migrations)

## Installation

Installing Gator via `go install`

The `go install` command compiles the Go package in your current directory and places the resulting executable binary into your Go binary directory (`$GOPATH/bin` or `$HOME/go/bin` by default).

1. Build and install the binary
   Run the following command from the root directory of your project:

    ```
    go install .
    ```

    This compiles your main package into an executable named after your module or project directory (typically gator).

2. Ensure your Go bin directory is in your `PATH`
   For your shell to recognize gator as a command from any directory, your Go binary directory must be in your system's PATH.

    To verify your Go binary path:

    ```
    go env GOPATH
    ```

    Binaries installed by `go install` live in the `bin` subdirectory under that path (usually `~/go/bin`).

    If running `gator` returns `command not found`, add the directory to your shell configuration file (such as `~/.bashrc`, `~/.zshrc`, or `~/.bash_profile`):

    ```
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```

    After updating the file, reload your shell:

    ```
    source ~/.bashrc   # or source ~/.zshrc
    ```

3. Run the CLI directly
   Once installed and present in your `PATH`, you no longer need `go run .`:

    ```
    gator feeds
    gator register alice
    gator addfeed "Tech News" "https://example.com/rss"
    ```
