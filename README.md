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

## Configuration

1. Set Up the Config File

    Gator reads its settings from a JSON config file located at ~/.gatorconfig.json in your home directory.

    Add this structure, adjusting the database URL to match your local PostgreSQL setup:

    ```
    {
    	"db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
    	"current_user_name": ""
    }
    ```

    `db_url`: the connection string Gator uses to talk to your PostgreSQL database.
    `current_user_name`: gets populated automatically once you register or log in as a user — you can leave it blank initially.

2. Set Up the Database Schema

    Gator relies on PostgreSQL tables defined via SQL migration files (found in `sql/schema/`). You'll use **Goose** to apply them.
    1. Make sure PostgreSQL is running and you've created a database (e.g., named `gator`):

    ```
    createdb gator
    ```

    2. Navigate to your schema folder:

    ```
    cd sql/schema
    ```

    3. Run the migrations "up" to create the necessary tables:

    ```
    goose postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" up
    ```

    This applies each versioned migration file in order, creating tables like `users` and `feeds`.

3. Run the Program / Use guide

    If you have not install or build the program, please do so by referring to the [Installation](#installation) Here are some commands that you can try:
    1. **Register a new user:**

    ```
    gator register alice
    ```

    Creates a new user and automatically sets them as your `current_user_name` in you .gatorconfig.json
    2. **Add a new RSS feed:**

    ```
    gator addfeed "Hacker News" "https://hnrss.org/newest"
    ```

    Saves the feed to the database, linked to whichever user is currently logged in.
    3. **List all feeds:**

    ```
    gator feeds
    ```

    Displays every feed stored in the database, including the feed's name, URL, and the user who added it.
    4. **Follows a feed:**

    ```
    gator follow https://www.wagslane.dev/index.xml
    ```

    User can follow an existing feed by providing it's URL. You may check exisiting feeds with `feeds` command.
    5. **Unfollow**

    ```
    gator unfollow
    ```

    User can unfollow a feed that they followed.
    6. **Aggregate**

    ```
    gator agg <time_between_requests>
    ```

    Aggregate allows continuous fetch on RSS feeds on a set interval (for example, 1m for every minute, or 30s for every 30 seconds), parses the XML, and stores new posts in the database. It runs indefinitely until you stop it manually (Ctrl+C).
    7. **Reset the database:**

    ```
    gator reset
    ```

    Clears existing users and feeds — handy for starting fresh during testing.
