# Go URL Shortener

A simple URL shortener service written in Go that uses an in-memory map as the database. It generates shortened URLs and allows redirection to the original URLs.

## Features

- Shorten a long URL using a simple API endpoint.
- Redirect to the original URL using the shortened URL.
- In-memory database using Go `map`.
- Simple JSON-based API for URL shortening.

## Getting Started

### Prerequisites

To run this project, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.16 or higher)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-url-shortener.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-url-shortener
   ```

3. Build the project:

   ```bash
   go build
   ```

### Usage

1. Start the server:

   ```bash
   ./go-url-shortener
   ```

   The server will start on `http://localhost:8080`.

2. To shorten a URL, send a POST request to `/shorten`:

   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"url":"https://example.com"}' http://localhost:8080/shorten
   ```

   This will return a JSON response with the shortened URL:

   ```json
   { "short_url": "http://localhost:8080/abc123" }
   ```

3. To use the shortened URL, simply open it in a web browser or use curl:

   ```bash
   curl -L http://localhost:8080/abc123
   ```

   This will redirect to the original URL.

## API Endpoints

- `POST /shorten`: Shorten a URL

  - Request body: `{"url": "https://example.com"}`
  - Response: `{"short_url": "http://localhost:8080/abc123"}`

- `GET /{shortCode}`: Redirect to the original URL

## Implementation Details

- The project uses the `net/http` package to handle HTTP requests.
- Shortened URLs are stored in an in-memory `map[string]string`.
- A simple random string generator is used to create short codes.

## Future Improvements

- Implement persistent storage (e.g., database) for URLs.
- Add user authentication and custom short codes.
- Implement rate limiting to prevent abuse.
- Add analytics for link clicks.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
