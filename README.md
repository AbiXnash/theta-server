# theta-server

Lightweight Go API server that exposes static JSON data for the Theta website.

## Features

- Echo v5 HTTP server
- API key protected `/api` endpoints
- One endpoint per JSON dataset in `data/`
- Configurable runtime via environment variables
- Health endpoint for uptime checks
- Simple developer workflow via `Makefile`

## Project structure

- `main.go`: app bootstrap and route wiring
- `internal/config`: environment configuration loader
- `internal/data`: JSON file loader
- `internal/http/handlers`: HTTP handlers
- `internal/http/middleware`: API key middleware
- `internal/http/router`: endpoint registration
- `data/`: JSON payload sources served by the API
- `API.md`: endpoint details and request examples

## Requirements

- Go `1.25.7` or compatible toolchain
- `make`

## Quick start

```bash
make run
```

Server starts on `http://localhost:8080` by default.

## Configuration

- `PORT` (default: `8080`)
- `THETA_API_KEY` (default: `somethings`)
- `THETA_DATA_DIR` (default: `data`)

Example:

```bash
PORT=9090 THETA_API_KEY=my-secret make run
```

## Development commands

- `make fmt`: format Go files
- `make vet`: run static analysis (`go vet`)
- `make test`: run tests
- `make build`: compile binary (`theta-server`)
- `make run`: run service locally
- `make check`: run fmt + vet + test + build
- `make clean`: remove built binary

## API

See [API.md](API.md) for full endpoint list and `curl` examples.

## Verification

```bash
make check
```

## License

No license file is currently included.
