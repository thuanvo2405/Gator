# Gator

`gator` is a CLI blog aggregator written in Go. It allows users to register, log in, follow RSS feeds, continuously fetch posts from followed feeds using a background worker, and browse saved posts directly in the terminal.

## Prerequisites

To run `gator`, you must have the following installed on your machine:

- [Go](https://go.dev/) (version 1.20 or later)
- [PostgreSQL](https://www.postgresql.org/)

Make sure your PostgreSQL server is running and you have a database created for the project.

## Installation

Install the `gator` executable directly using `go install`:

```bash
go install [github.com/](https://github.com/)<your-username>/gator@latest
```

Configuration Setup
gator requires a JSON configuration file located at ~/.gatorconfig.json in your home directory.

Create the file with your PostgreSQL connection string and active user context:

JSON
{
"db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
"current_user_name": ""
}
Replace postgres:postgres@localhost:5432/gator with your local PostgreSQL credentials and database name.

Usage & Commands
Run commands using the compiled binary:

Bash
gator <command> [arguments]
Key Commands
gator register <name>

Register a new user in the database and set them as the active user in .gatorconfig.json.

gator login <name>

Switch the currently active user context to an existing user.

gator addfeed <name> <url>

Add a new RSS feed to the database and automatically subscribe the current user to it.

gator feeds

List all RSS feeds registered in the system along with their creators.

gator follow <url>

Subscribe the active user to an existing RSS feed by its URL.

gator following

Display all RSS feeds followed by the currently logged-in user.

gator agg <time_between_reqs>

Start the continuous aggregator process to fetch posts from feeds in the background (e.g., gator agg 1m).

gator browse [limit]

View the latest blog posts saved in your database for feeds you follow (defaults to 2 posts).
