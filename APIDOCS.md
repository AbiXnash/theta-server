# API Docs

Base URL: `http://localhost:8080`

## Authentication

All `/api/*` endpoints require:

- Header: `X-ABX-THETA-KEY: <your-api-key>`

Example:

```bash
curl -H "X-ABX-THETA-KEY: my-secret" http://localhost:8080/api/team
```

## Public Endpoint

### `GET /health`

Response:

```json
{"status":"ok"}
```

## Protected Endpoints

### `GET /api/team`
Serves `data/team.json`

### `GET /api/config`
Serves `data/config.json`

### `GET /api/events`
Serves `data/events.json`

### `GET /api/content`
Serves `data/content.json`

### `GET /api/sponsors`
Serves `data/sponsors.json`

### `GET /api/manifest`
Serves `data/manifest.json`

## Error Responses

- `401 Unauthorized`: invalid or missing API key
- `500 Internal Server Error`: server failed to load JSON file

## Environment Variables

- `PORT` (default `8080`)
- `THETA_API_KEY` (required, no default)
- `THETA_DATA_DIR` (default `data`)
