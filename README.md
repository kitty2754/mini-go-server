# Mini Go Wiki Server

A lightweight, minimalist wiki application built with Go. Create, edit, and view wiki pages with a simple and elegant interface.

## Features

- **Create & Edit Pages**: Easily create new wiki pages and edit existing ones
- **Simple File Storage**: Pages are stored as plain text files in the `data/` directory
- **Web Interface**: Clean, teal-themed web interface for viewing and editing
- **RESTful API**: HTTP endpoints for all operations
- **Health Check**: Built-in health check endpoint for monitoring

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

1. Clone the repository:
```bash
git clone https://github.com/kitty2754/mini-go-server.git
cd mini-go-server
```

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run main.go wiki.go types/pageTypes.go
```

The server will start on `http://localhost:8080`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check endpoint |
| GET | `/view/{title}` | View a wiki page |
| GET | `/edit/{title}` | Edit page for a wiki page |
| POST | `/save/{title}` | Save a wiki page |

### Examples

**View a page:**
```
GET http://localhost:8080/view/MyPage
```

**Edit a page:**
```
GET http://localhost:8080/edit/MyPage
```

**Save a page:**
```
POST http://localhost:8080/save/MyPage
Body: body=<content>
```

## Project Structure

```
.
├── main.go              # Main application and HTTP handlers
├── wiki.go              # Page model and file operations
├── types/
│   └── pageTypes.go    # Page struct definition
├── template/
│   ├── view.html       # Page view template
│   └── edit.html       # Page edit template
├── data/                # Directory for storing wiki pages
├── go.mod              # Go module file
└── README.md           # This file
```

## How It Works

1. **Page Storage**: Wiki pages are stored as individual `.txt` files in the `data/` directory
2. **Page Structure**: Each page contains a title and body content
3. **Editing**: Users can access the edit interface to create or modify page content
4. **Viewing**: Pages are rendered with a teal background for a pleasant user experience

## Configuration

- **Server Port**: The server runs on port `8080` (configured in `main.go`)
- **Data Directory**: Pages are stored in the `data/` folder
- **File Format**: Pages are saved as `.txt` files with the page title as the filename

## Development

To build the project:
```bash
go build -o hello main.go wiki.go types/pageTypes.go
```

To run tests (if any):
```bash
go test ./...
```

## Styling

The wiki pages feature a teal-colored background for an attractive user interface. The styling is defined in `template/view.html`.

## Future Enhancements

- [ ] Add page deletion functionality
- [ ] Implement user authentication
- [ ] Add markdown support
- [ ] Add page versioning/history
- [ ] Implement search functionality

## License

MIT License - Feel free to use this project for personal or educational purposes.

## Contributing

Contributions are welcome! Feel free to submit pull requests or open issues for bugs and feature requests.