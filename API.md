# Theta Server API

Base URL (local): `http://localhost:8080`

## Authentication

All `/api/*` routes require header:

- `X-ABX-THETA-KEY: <your key>`

Example:

```bash
curl -H "X-ABX-THETA-KEY: my-secret" http://localhost:8080/api/team
```

## Endpoints

Protected endpoints:

- `GET /api/team` -> `data/team.json`
- `GET /api/config` -> `data/config.json`
- `GET /api/events` -> `data/events.json`
- `GET /api/content` -> `data/content.json`
- `GET /api/sponsors` -> `data/sponsors.json`
- `GET /api/manifest` -> `data/manifest.json`

Public endpoints:

- `GET /health` -> `{"status":"ok"}`

## Errors

- `401 Unauthorized`: missing or invalid API key
- `500 Internal Server Error`: file read/load failure

## Environment variables

- `PORT` (default: `8080`)
- `THETA_API_KEY` (required, no default)
- `THETA_DATA_DIR` (default: `data`)
