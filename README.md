# URL Shortener

A simple URL Shortener service built with Go using the standard `net/http` package. This project allows users to create short URLs that redirect to long URLs, similar to services like Bitly or TinyURL.

## Features

* Create shortened URLs
* Redirect short URLs to original URLs
* RESTful HTTP endpoints
* Built using Go's standard library (`net/http`)
* Lightweight and dependency-free
* In-memory URL storage (can be extended to use a database)

## Project Structure

```text
urlshortener/
├── main.go
├── models.go
|── handles.go
│
├── go.mod
└── README.md
```

## Installation

### Prerequisites

* Go 1.22 or later

### Clone the Repository

```bash
git clone https://github.com/michaelbag8/urlshortener.git
cd urlshortener
```

### Run the Application

```bash
go run main.go
```

The server will start on:

```text
http://localhost:8080
```

## API Endpoints

### Create a Short URL

**POST** `/shorten`

#### Request

```json
{
  "url": "https://www.example.com"
}
```

#### Response

```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

### Redirect to Original URL

**GET** `/{shortCode}`

Example:

```text
GET /abc123
```

Response:

```http
302 Found
Location: https://www.example.com
```

## Example Usage

Create a short URL:

```bash
curl -X POST http://localhost:8080/shorten \
-H "Content-Type: application/json" \
-d '{"url":"https://golang.org"}'
```

Visit the generated short URL:

```text
http://localhost:8080/abc123
```

The application redirects the user to the original URL using:

```go
http.Redirect(w, r, originalURL, http.StatusFound)
```

## How It Works

1. User submits a long URL.
2. The server generates a unique short code.
3. The URL mapping is stored in memory.
4. The server returns the shortened URL.
5. When a user visits the short URL, the server looks up the original URL and redirects the request.

## Future Improvements

* Persistent database storage (PostgreSQL, MySQL, SQLite)
* URL expiration
* Custom aliases
* Analytics and click tracking
* User authentication
* Rate limiting
* Docker support

## Technologies Used

* Go
* net/http
* encoding/json
* Standard Library

