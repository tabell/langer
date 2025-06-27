# langer
Language Learning app

## Backend Go Server

The `backend-go` directory contains a small Go web server with several stub endpoints:

- `/ping` returns a simple JSON response for health checks.
- `/lookup` looks up a word or phrase, returning placeholder translation and knowledge information.
- `/import` accepts new articles or audio files for processing.
- `/content` serves previously imported content.

### Building

```bash
cd backend-go
go build -o langer-server
```

### Running

```bash
./langer-server
```

The server listens on port `8080` by default and serves files from the `static/` directory.
Open `http://localhost:8080/` after running the server to see the first available content.
