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
