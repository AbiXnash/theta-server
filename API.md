# Theta Server API

All endpoints are mounted under `/api` and require this header:

- `X-ABX-THETA-KEY: <your key>`

Default key is `somethings` unless `THETA_API_KEY` is set.

## Endpoints

- `GET /api/team` -> `data/team.json`
- `GET /api/config` -> `data/config.json`
- `GET /api/events` -> `data/events.json`
- `GET /api/content` -> `data/content.json`
- `GET /api/sponsors` -> `data/sponsors.json`
- `GET /api/manifest` -> `data/manifest.json`

## Environment variables

- `PORT` (default: `8080`)
- `THETA_API_KEY` (default: `somethings`)
- `THETA_DATA_DIR` (default: `data`)

- `GET /health` -> service health check
