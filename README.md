# TDR3

Work in progress! Super simple time tracking application. Written in Go and Vue and uses SQLite as database.

## Install

```
go mod download
```

```
cd frontend
cp .env.example .env
npm install
```

## Start

To start backend:

`air` or `go run cmd/web/main.go`

To start frontend:

`cd frontend && npm run serve`

## ENV variables

`PRODUCTION = true` for hiding sql queries in the log
